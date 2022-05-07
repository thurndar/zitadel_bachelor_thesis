package user

import (
	"context"
	"encoding/json"
	"time"

	"github.com/caos/zitadel/internal/domain"
	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/repository"
)

const (
	machineKeyEventPrefix      = machineEventPrefix + "key."
	MachineKeyAddedEventType   = machineKeyEventPrefix + "added"
	MachineKeyRemovedEventType = machineKeyEventPrefix + "removed"
)

type MachineKeyAddedEvent struct {
	eventstore.BaseEvent

	KeyID          string              `json:"keyId,omitempty"`
	KeyType        domain.AuthNKeyType `json:"type,omitempty"`
	ExpirationDate time.Time           `json:"expirationDate,omitempty"`
	PublicKey      []byte              `json:"publicKey,omitempty"`
}

func (e *MachineKeyAddedEvent) Data() interface{} {
	return e
}

func (e *MachineKeyAddedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewMachineKeyAddedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	keyID string,
	keyType domain.AuthNKeyType,
	expirationDate time.Time,
	publicKey []byte,
) *MachineKeyAddedEvent {
	return &MachineKeyAddedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			MachineKeyAddedEventType,
		),
		KeyID:          keyID,
		KeyType:        keyType,
		ExpirationDate: expirationDate,
		PublicKey:      publicKey,
	}
}

func MachineKeyAddedEventMapper(event *repository.Event) (eventstore.Event, error) {
	machineKeyAdded := &MachineKeyAddedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}
	err := json.Unmarshal(event.Data, machineKeyAdded)
	if err != nil {
		//first events had wrong payload.
		// the keys were removed later, that's why we ignore them here.
		if unwrapErr, ok := err.(*json.UnmarshalTypeError); ok && unwrapErr.Field == "publicKey" {
			return machineKeyAdded, nil
		}
		return nil, errors.ThrowInternal(err, "USER-p0ovS", "unable to unmarshal machine key added")
	}

	machineKeyAdded.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return machineKeyAdded, nil
}

type MachineKeyRemovedEvent struct {
	eventstore.BaseEvent

	KeyID string `json:"keyId,omitempty"`
}

func (e *MachineKeyRemovedEvent) Data() interface{} {
	return e
}

func (e *MachineKeyRemovedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewMachineKeyRemovedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	keyID string,
) *MachineKeyRemovedEvent {
	return &MachineKeyRemovedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			MachineKeyRemovedEventType,
		),
		KeyID: keyID,
	}
}

func MachineKeyRemovedEventMapper(event *repository.Event) (eventstore.Event, error) {
	machineRemoved := &MachineKeyRemovedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}
	err := json.Unmarshal(event.Data, machineRemoved)
	if err != nil {
		return nil, errors.ThrowInternal(err, "USER-5Gm9s", "unable to unmarshal machine key removed")
	}
	machineRemoved.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return machineRemoved, nil
}
