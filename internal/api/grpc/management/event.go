package management

import (
	"context"

	"github.com/caos/zitadel/internal/api/authz"
	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	mgmt_pb "github.com/caos/zitadel/pkg/grpc/management"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Server) ListEvents(ctx context.Context, req *mgmt_pb.ListEventsRequest) (*mgmt_pb.ListEventsResponse, error) {
	if err := checkPerimissionsOnListEventRequest(ctx, req); err != nil {
		return nil, err
	}

	events, err := s.es.Filter(ctx, eventsRequestToBuilder(req))
	if err != nil {
		return nil, err
	}

	return &mgmt_pb.ListEventsResponse{
		Events: eventsToAPI(events),
	}, nil
}

func eventsToAPI(events []eventstore.Event) []*mgmt_pb.Event {
	response := make([]*mgmt_pb.Event, len(events))
	for i, event := range events {
		response[i] = eventToAPI(event)
	}

	return response
}

func eventToAPI(event eventstore.Event) *mgmt_pb.Event {
	return &mgmt_pb.Event{
		Sequence:     event.Sequence(),
		CreationDate: timestamppb.New(event.CreationDate()),
		Type:         string(event.Type()),
		Editor:       event.EditorUser(),
		Payload:      event.DataAsBytes(),
		Aggregate: &mgmt_pb.Aggregate{
			Id:            event.Aggregate().ID,
			Type:          string(event.Aggregate().Type),
			Version:       string(event.Aggregate().Version),
			ResourceOwner: event.Aggregate().ResourceOwner,
		},
	}
}

func eventsRequestToBuilder(req *mgmt_pb.ListEventsRequest) *eventstore.SearchQueryBuilder {
	builder := eventstore.NewSearchQueryBuilder(eventstore.ColumnsEvent)
	builder = builder.Limit(uint64(req.Limit))

	builder = builder.OrderAsc()
	if req.Desc {
		builder = builder.OrderAsc()
	}

	for _, aggregate := range req.Aggregates {
		aggregateTypes := make([]eventstore.AggregateType, len(aggregate.Types))
		for i, typ := range aggregate.Types {
			aggregateTypes[i] = eventstore.AggregateType(typ)
		}
		query := builder.AddQuery().
			AggregateTypes(aggregateTypes...).
			AggregateIDs(aggregate.Ids...).
			SequenceGreater(aggregate.SequenceGreater).
			SequenceLess(aggregate.SequenceLess)

		if aggregate.Event != nil {
			eventTypes := make([]eventstore.EventType, len(aggregate.Event.Types))
			for i, typ := range aggregate.Event.Types {
				eventTypes[i] = eventstore.EventType(typ)
			}
			query.EventTypes(eventTypes...)
		}

		builder = query.Builder()
	}

	return builder
}

func checkPerimissionsOnListEventRequest(ctx context.Context, req *mgmt_pb.ListEventsRequest) error {
	permissions := authz.GetAllPermissionsFromCtx(ctx)
	for _, aggregate := range req.Aggregates {
		if err := checkPermissionsOnAggregate(permissions, aggregate); err != nil {
			return err
		}
	}

	return nil
}

func checkPermissionsOnAggregate(permissions []string, aggregate *mgmt_pb.AggregateQuery) error {
	for _, typ := range aggregate.Types {
		if err := containsReadRole(permissions, typ); err != nil {
			return err
		}
	}

	return nil
}

func containsReadRole(permissions []string, typ string) error {
	for _, permission := range permissions {
		if permission == typ+".read" {
			return nil
		}
	}

	return errors.ThrowInvalidArgument(nil, "MANAG-3gp7m", "Errors.Insufficient.rights")
}
