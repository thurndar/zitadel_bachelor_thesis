package eventstore

import (
	"encoding/json"
	"log"
	"sync"

	v1 "github.com/caos/zitadel/internal/eventstore/v1"
	"github.com/caos/zitadel/internal/eventstore/v1/models"
	"github.com/caos/zitadel/internal/telemetry/tracing"
	"github.com/nats-io/nats.go"
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

//SubscribeAggregates subscribes for all events on the given aggregates
func SubscribeAggregates(eventQueue chan Event, aggregates ...AggregateType) *Subscription {
	// types := make(map[AggregateType][]EventType, len(aggregates))
	// for _, aggregate := range aggregates {
	// 	types[aggregate] = nil
	// }
	// sub := &Subscription{
	// 	Events: eventQueue,
	// 	types:  types,
	// }

	// subsMutext.Lock()
	// defer subsMutext.Unlock()

	// for _, aggregate := range aggregates {
	// 	subscriptions[aggregate] = append(subscriptions[aggregate], sub)
	// }

	// return sub

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

	// subsMutext.Lock()
	// defer subsMutext.Unlock()

	// for _, aggregate := range aggregates {
	// 	subscriptions[aggregate] = append(subscriptions[aggregate], sub)
	// }
	for aggregate := range types {
		log.Printf("Try to subscribe to %s", aggregate)
		_, err := nc.Subscribe(string(aggregate)+".>", func(msg *nats.Msg) {
			log.Printf("Entering the function call of event: %s", aggregate)
			_, subscribeEventTypeSpan := tracing.NewNamedSpan(context.TODO(), "subscribeEventTypeSpan")
			var err error
			defer func() { subscribeEventTypeSpan.EndWithError(err) }()

			// nats msg to zitadel event
			var event = eventStructs[EventType(msg.Subject)]()
			// tracing.SetLabel(subscribeEventTypeSpan, "event.type", string(event.Type()))
			err = json.Unmarshal(msg.Data, event)
			if err != nil {
				log.Printf("err in unmarshal: %v", err)
			}

			sub.Events <- event
			subscribeEventTypeSpan.End()
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
		// subs, ok := subscriptions[event.Aggregate().Type]
		// if !ok {
		// 	continue
		// }
		// for _, sub := range subs {
		// 	eventTypes := sub.types[event.Aggregate().Type]
		// 	//subscription for all events
		// 	if len(eventTypes) == 0 {
		// 		sub.Events <- event
		// 		continue
		// 	}
		// 	//subscription for certain events
		// 	for _, eventType := range eventTypes {
		// 		if event.Type() == eventType {
		// 			sub.Events <- event
		// 			break
		// 		}
		// 	}
		// }

		// span is missing labels, we don't know which event this mothertrucker is
		// eventtype
		_, notifySpan := tracing.NewNamedSpan(ctx, "notify")
		var err error
		defer func() { notifySpan.EndWithError(err) }()
		// tracing.SetLabel(notifySpan, "event.type", string(event.Type()))

		msg, _ := json.Marshal(event)
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
