package eventstore

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"sync"

	v1 "github.com/caos/zitadel/internal/eventstore/v1"
	"github.com/caos/zitadel/internal/eventstore/v1/models"
	"github.com/caos/zitadel/internal/telemetry/tracing"
	"github.com/nats-io/nats.go"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"
)

var (
	subscriptions = map[AggregateType][]*Subscription{}
	subsMutext    sync.Mutex
)

type Subscription struct {
	Events chan Event
	types  map[AggregateType][]EventType
}

type notifyEvent struct {
	SpanCtx trace.SpanContext
	Event   Event
}

func (e *notifyEvent) UnmarshalJSON(data []byte) (err error) {
	fields := map[string]interface{}{}
	if err := json.Unmarshal(data, &fields); err != nil {
		return err
	}
	spanContext := fields["SpanCtx"].(map[string]interface{})
	spanConfig := trace.SpanContextConfig{}

	if traceID, ok := spanContext["TraceID"].(string); ok {
		spanConfig.TraceID, err = trace.TraceIDFromHex(traceID)
		if err != nil {
			return err
		}
	}
	if spanID, ok := spanContext["SpanID"].(string); ok {
		if spanConfig.SpanID, err = trace.SpanIDFromHex(spanID); err != nil {
			return err
		}
	}
	if traceFlags, ok := spanContext["TraceFlags"].(string); ok {
		if _, err := hex.Decode([]byte{byte(spanConfig.TraceFlags)}[:], []byte(traceFlags)); err != nil {
			return err
		}
	}

	e.SpanCtx = trace.NewSpanContext(spanConfig)

	eventJSON, err := json.Marshal(fields["Event"])
	if err != nil {
		return err
	}
	if err := json.Unmarshal(eventJSON, e.Event); err != nil {
		return err
	}

	return nil
}

//SubscribeAggregates subscribes for all events on the given aggregates
func SubscribeAggregates(eventQueue chan Event, aggregates ...AggregateType) *Subscription {
	types := make(map[AggregateType][]EventType, len(aggregates))
	for _, aggregate := range aggregates {
		types[aggregate] = nil
	}

	return SubscribeEventTypes(eventQueue, types)
}

//SubscribeEventTypes subscribes for the given event types
// if no event types are provided the subscription is for all events of the aggregate
// Subscriber but in wrong package :-)
func SubscribeEventTypes(eventQueue chan Event, types map[AggregateType][]EventType) *Subscription {
	// aggregates := make([]AggregateType, len(types))
	sub := &Subscription{
		Events: eventQueue,
		types:  types,
	}

	for aggregate := range types {
		_, err := nc.Subscribe(string(aggregate)+".>", func(msg *nats.Msg) {
			// nats msg to zitadel event
			event := eventStructs[EventType(msg.Subject)]()

			e := &notifyEvent{
				Event: event,
			}
			// tracing.SetLabel(subscribeEventTypeSpan, "event.type", string(event.Type()))
			err := json.Unmarshal(msg.Data, e)
			if err != nil {
				log.Printf("unable to unmarshal: %v\n", err)
			}

			// ctx := trace.ContextWithRemoteSpanContext(context.TODO(), e.SpanCtx)
			_, span := tracing.NewNamedSpan(context.TODO(), "subscribeEventTypeSpan", trace.WithLinks(trace.Link{SpanContext: e.SpanCtx}))
			defer func() { span.EndWithError(err) }()

			sub.Events <- event
		})
		if err != nil {
			log.Println(err)
		}
	}

	return sub
}

// actually publisher
func notify(ctx context.Context, events []Event) {
	go v1.Notify(MapEventsToV1Events(events))
	subsMutext.Lock()
	defer subsMutext.Unlock()
	for _, event := range events {
		// span is missing labels, we don't know which event this mothertrucker is
		// eventtype
		_, notifySpan := tracing.NewNamedSpan(ctx, "notify")
		var err error
		defer func() { notifySpan.EndWithError(err) }()
		// tracing.SetLabel(notifySpan, "event.type", string(event.Type()))
		msg, _ := json.Marshal(&notifyEvent{
			SpanCtx: notifySpan.OTELSpan().SpanContext(),
			Event:   event,
		})

		err = nc.Publish(string(event.Type()), msg)
		if err != nil {
			log.Println(err)
		}

		notifySpan.End()
	}
}

func (s *Subscription) Unsubscribe() {
	subsMutext.Lock()
	defer subsMutext.Unlock()
	for aggregate := range s.types {
		subs, ok := subscriptions[aggregate]
		if !ok {
			continue
		}
		for i := len(subs) - 1; i >= 0; i-- {
			if subs[i] == s {
				subs[i] = subs[len(subs)-1]
				subs[len(subs)-1] = nil
				subs = subs[:len(subs)-1]
			}
		}
	}
	_, ok := <-s.Events
	if ok {
		close(s.Events)
	}
}

func MapEventsToV1Events(events []Event) []*models.Event {
	v1Events := make([]*models.Event, len(events))
	for i, event := range events {
		v1Events[i] = mapEventToV1Event(event)
	}
	return v1Events
}

func mapEventToV1Event(event Event) *models.Event {
	return &models.Event{
		Sequence:      event.Sequence(),
		CreationDate:  event.CreationDate(),
		Type:          models.EventType(event.Type()),
		AggregateType: models.AggregateType(event.Aggregate().Type),
		AggregateID:   event.Aggregate().ID,
		ResourceOwner: event.Aggregate().ResourceOwner,
		EditorService: event.EditorService(),
		EditorUser:    event.EditorUser(),
		Data:          event.DataAsBytes(),
	}
}

func RegisterEventType(typ EventType, creator func() Event) {
	eventStructsMu.Lock()
	if _, ok := eventStructs[typ]; !ok {
		eventStructs[typ] = creator
	}
	eventStructsMu.Unlock()
}

var (
	eventStructs   = map[EventType]func() Event{}
	eventStructsMu = sync.Mutex{}
)
