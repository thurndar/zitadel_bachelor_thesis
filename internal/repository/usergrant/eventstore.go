package usergrant

import (
	"github.com/caos/zitadel/internal/eventstore"
)

func RegisterEventMappers(es *eventstore.Eventstore) {
	es.RegisterFilterEventMapper(UserGrantAddedType, UserGrantAddedEventMapper).
		RegisterFilterEventMapper(UserGrantChangedType, UserGrantChangedEventMapper).
		RegisterFilterEventMapper(UserGrantCascadeChangedType, UserGrantCascadeChangedEventMapper).
		RegisterFilterEventMapper(UserGrantRemovedType, UserGrantRemovedEventMapper).
		RegisterFilterEventMapper(UserGrantCascadeRemovedType, UserGrantCascadeRemovedEventMapper).
		RegisterFilterEventMapper(UserGrantDeactivatedType, UserGrantDeactivatedEventMapper).
		RegisterFilterEventMapper(UserGrantReactivatedType, UserGrantReactivatedEventMapper)

	eventstore.RegisterEventType(UserGrantAddedType, func() eventstore.Event {
		return new(UserGrantAddedEvent)
	})
	eventstore.RegisterEventType(UserGrantChangedType, func() eventstore.Event {
		return new(UserGrantChangedEvent)
	})
	eventstore.RegisterEventType(UserGrantCascadeChangedType, func() eventstore.Event {
		return new(UserGrantCascadeChangedEvent)
	})
	eventstore.RegisterEventType(UserGrantRemovedType, func() eventstore.Event {
		return new(UserGrantRemovedEvent)
	})
	eventstore.RegisterEventType(UserGrantCascadeRemovedType, func() eventstore.Event {
		return new(UserGrantCascadeRemovedEvent)
	})
	eventstore.RegisterEventType(UserGrantDeactivatedType, func() eventstore.Event {
		return new(UserGrantDeactivatedEvent)
	})
	eventstore.RegisterEventType(UserGrantReactivatedType, func() eventstore.Event {
		return new(UserGrantReactivatedEvent)
	})
}
