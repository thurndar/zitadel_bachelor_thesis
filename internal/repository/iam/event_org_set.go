package iam

import (
	"context"
	"encoding/json"

	"github.com/caos/zitadel/internal/eventstore"

	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore/repository"
)

const (
	GlobalOrgSetEventType eventstore.EventType = "iam.global.org.set"
)

type GlobalOrgSetEvent struct {
	eventstore.BaseEvent

	OrgID string `json:"globalOrgId"`
}

func (e *GlobalOrgSetEvent) Data() interface{} {
	return e
}

func (e *GlobalOrgSetEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewGlobalOrgSetEventEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	orgID string,
) *GlobalOrgSetEvent {
	return &GlobalOrgSetEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			GlobalOrgSetEventType,
		),
		OrgID: orgID,
	}
}

func GlobalOrgSetMapper(event *repository.Event) (eventstore.Event, error) {
	e := &GlobalOrgSetEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}
	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "IAM-cdFZH", "unable to unmarshal global org set")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}
