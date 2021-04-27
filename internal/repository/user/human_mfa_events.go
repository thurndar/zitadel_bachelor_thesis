package user

import (
	"context"
	"github.com/caos/zitadel/internal/eventstore"

	"github.com/caos/zitadel/internal/eventstore/repository"
)

const (
	mfaEventPrefix          = humanEventPrefix + "mfa."
	HumanMFAInitSkippedType = mfaEventPrefix + "init.skipped"
)

type HumanMFAInitSkippedEvent struct {
	eventstore.BaseEvent `json:"-"`
}

func (e *HumanMFAInitSkippedEvent) Data() interface{} {
	return e
}

func (e *HumanMFAInitSkippedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func (e *HumanMFAInitSkippedEvent) Assets() []*eventstore.Asset {
	return nil
}

func NewHumanMFAInitSkippedEvent(ctx context.Context, aggregate *eventstore.Aggregate) *HumanMFAInitSkippedEvent {
	return &HumanMFAInitSkippedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			HumanMFAInitSkippedType,
		),
	}
}

func HumanMFAInitSkippedEventMapper(event *repository.Event) (eventstore.EventReader, error) {
	return &HumanMFAInitSkippedEvent{
		BaseEvent: *eventstore.BaseEventFromRepo(event),
	}, nil
}
