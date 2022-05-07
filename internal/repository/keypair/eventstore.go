package keypair

import (
	"github.com/caos/zitadel/internal/eventstore"
)

func RegisterEventMappers(es *eventstore.Eventstore) {
	es.RegisterFilterEventMapper(AddedEventType, AddedEventMapper)

	eventstore.RegisterEventType(AddedEventType, func() eventstore.Event { return new(AddedEvent) })
}
