package project

import (
	"context"
	"encoding/json"
	"time"

	"github.com/caos/zitadel/internal/eventstore"

	"github.com/caos/zitadel/internal/domain"
	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore/repository"
)

const (
	applicationKeyEventPrefix      = applicationEventTypePrefix + "oidc.key."
	ApplicationKeyAddedEventType   = applicationKeyEventPrefix + "added"
	ApplicationKeyRemovedEventType = applicationKeyEventPrefix + "removed"
)

type ApplicationKeyAddedEvent struct {
	eventstore.BaseEvent

	AppID          string              `json:"applicationId"`
	ClientID       string              `json:"clientId,omitempty"`
	KeyID          string              `json:"keyId,omitempty"`
	KeyType        domain.AuthNKeyType `json:"type,omitempty"`
	ExpirationDate time.Time           `json:"expirationDate,omitempty"`
	PublicKey      []byte              `json:"publicKey,omitempty"`
}

func (e *ApplicationKeyAddedEvent) Data() interface{} {
	return e
}

func (e *ApplicationKeyAddedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewApplicationKeyAddedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	appID,
	clientID,
	keyID string,
	keyType domain.AuthNKeyType,
	expirationDate time.Time,
	publicKey []byte,
) *ApplicationKeyAddedEvent {
	return &ApplicationKeyAddedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			ApplicationKeyAddedEventType,
		),
		AppID:          appID,
		ClientID:       clientID,
		KeyID:          keyID,
		KeyType:        keyType,
		ExpirationDate: expirationDate,
		PublicKey:      publicKey,
	}
}

func ApplicationKeyAddedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &ApplicationKeyAddedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "API-BFd15", "unable to unmarshal api config")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}

type ApplicationKeyRemovedEvent struct {
	eventstore.BaseEvent

	KeyID string `json:"keyId,omitempty"`
}

func (e *ApplicationKeyRemovedEvent) Data() interface{} {
	return e
}

func (e *ApplicationKeyRemovedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewApplicationKeyRemovedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	keyID string,
) *ApplicationKeyRemovedEvent {
	return &ApplicationKeyRemovedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			ApplicationKeyRemovedEventType,
		),
		KeyID: keyID,
	}
}

func ApplicationKeyRemovedEventMapper(event *repository.Event) (eventstore.Event, error) {
	applicationKeyRemoved := &ApplicationKeyRemovedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}
	err := json.Unmarshal(event.Data, applicationKeyRemoved)
	if err != nil {
		return nil, errors.ThrowInternal(err, "USER-5Gm9s", "unable to unmarshal application key removed")
	}
	applicationKeyRemoved.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return applicationKeyRemoved, nil
}
