package action

import "github.com/caos/zitadel/internal/eventstore"

func RegisterEventMappers(es *eventstore.Eventstore) {
	es.RegisterFilterEventMapper(AddedEventType, AddedEventMapper).
		RegisterFilterEventMapper(ChangedEventType, ChangedEventMapper).
		RegisterFilterEventMapper(DeactivatedEventType, DeactivatedEventMapper).
		RegisterFilterEventMapper(ReactivatedEventType, ReactivatedEventMapper).
		RegisterFilterEventMapper(RemovedEventType, RemovedEventMapper)

	eventstore.RegisterEventType(AddedEventType, func() eventstore.Event { return new(AddedEvent) })
	eventstore.RegisterEventType(ChangedEventType, func() eventstore.Event { return new(ChangedEvent) })
	eventstore.RegisterEventType(DeactivatedEventType, func() eventstore.Event { return new(DeactivatedEvent) })
	eventstore.RegisterEventType(ReactivatedEventType, func() eventstore.Event { return new(ReactivatedEvent) })
	eventstore.RegisterEventType(RemovedEventType, func() eventstore.Event { return new(RemovedEvent) })
}
