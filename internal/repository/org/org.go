package org

import (
	"context"
	"encoding/json"

	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/repository"
)

const (
	uniqueOrgname           = "org_name"
	OrgAddedEventType       = orgEventTypePrefix + "added"
	OrgChangedEventType     = orgEventTypePrefix + "changed"
	OrgDeactivatedEventType = orgEventTypePrefix + "deactivated"
	OrgReactivatedEventType = orgEventTypePrefix + "reactivated"
	OrgRemovedEventType     = orgEventTypePrefix + "removed"
)

func NewAddOrgNameUniqueConstraint(orgName string) *eventstore.EventUniqueConstraint {
	return eventstore.NewAddEventUniqueConstraint(
		uniqueOrgname,
		orgName,
		"Errors.Org.AlreadyExists")
}

func NewRemoveOrgNameUniqueConstraint(orgName string) *eventstore.EventUniqueConstraint {
	return eventstore.NewRemoveEventUniqueConstraint(
		uniqueOrgname,
		orgName)
}

type OrgAddedEvent struct {
	eventstore.BaseEvent

	Name string `json:"name,omitempty"`
}

func (e *OrgAddedEvent) Data() interface{} {
	return e
}

func (e *OrgAddedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return []*eventstore.EventUniqueConstraint{NewAddOrgNameUniqueConstraint(e.Name)}
}

func NewOrgAddedEvent(ctx context.Context, aggregate *eventstore.Aggregate, name string) *OrgAddedEvent {
	return &OrgAddedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			OrgAddedEventType,
		),
		Name: name,
	}
}

func OrgAddedEventMapper(event *repository.Event) (eventstore.Event, error) {
	orgAdded := &OrgAddedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}
	err := json.Unmarshal(event.Data, orgAdded)
	if err != nil {
		return nil, errors.ThrowInternal(err, "ORG-Bren2", "unable to unmarshal org added")
	}
	orgAdded.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return orgAdded, nil
}

type OrgChangedEvent struct {
	eventstore.BaseEvent

	Name    string `json:"name,omitempty"`
	oldName string `json:"-"`
}

func (e *OrgChangedEvent) Data() interface{} {
	return e
}

func (e *OrgChangedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return []*eventstore.EventUniqueConstraint{
		NewRemoveOrgNameUniqueConstraint(e.oldName),
		NewAddOrgNameUniqueConstraint(e.Name),
	}
}

func NewOrgChangedEvent(ctx context.Context, aggregate *eventstore.Aggregate, oldName, newName string) *OrgChangedEvent {
	return &OrgChangedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			OrgChangedEventType,
		),
		Name:    newName,
		oldName: oldName,
	}
}

func OrgChangedEventMapper(event *repository.Event) (eventstore.Event, error) {
	orgChanged := &OrgChangedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}
	err := json.Unmarshal(event.Data, orgChanged)
	if err != nil {
		return nil, errors.ThrowInternal(err, "ORG-Bren2", "unable to unmarshal org added")
	}
	orgChanged.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return orgChanged, nil
}

type OrgDeactivatedEvent struct {
	eventstore.BaseEvent
}

func (e *OrgDeactivatedEvent) Data() interface{} {
	return e
}

func (e *OrgDeactivatedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewOrgDeactivatedEvent(ctx context.Context, aggregate *eventstore.Aggregate) *OrgDeactivatedEvent {
	return &OrgDeactivatedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			OrgDeactivatedEventType,
		),
	}
}

func OrgDeactivatedEventMapper(event *repository.Event) (eventstore.Event, error) {
	return &OrgDeactivatedEvent{
		BaseEvent: *eventstore.BaseEventFromRepo(event),
	}, nil
}

type OrgReactivatedEvent struct {
	eventstore.BaseEvent
}

func (e *OrgReactivatedEvent) Data() interface{} {
	return e
}

func (e *OrgReactivatedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewOrgReactivatedEvent(ctx context.Context, aggregate *eventstore.Aggregate) *OrgReactivatedEvent {
	return &OrgReactivatedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			OrgReactivatedEventType,
		),
	}
}

func OrgReactivatedEventMapper(event *repository.Event) (eventstore.Event, error) {
	return &OrgReactivatedEvent{
		BaseEvent: *eventstore.BaseEventFromRepo(event),
	}, nil
}

type OrgRemovedEvent struct {
	eventstore.BaseEvent
	name string
}

func (e *OrgRemovedEvent) Data() interface{} {
	return e
}

func (e *OrgRemovedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return []*eventstore.EventUniqueConstraint{NewRemoveOrgNameUniqueConstraint(e.name)}
}

func NewOrgRemovedEvent(ctx context.Context, aggregate *eventstore.Aggregate, name string) *OrgRemovedEvent {
	return &OrgRemovedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			OrgRemovedEventType,
		),
		name: name,
	}
}

func OrgRemovedEventMapper(event *repository.Event) (eventstore.Event, error) {
	orgChanged := &OrgRemovedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}
	err := json.Unmarshal(event.Data, orgChanged)
	if err != nil {
		return nil, errors.ThrowInternal(err, "ORG-DAfbs", "unable to unmarshal org deactivated")
	}
	orgChanged.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return orgChanged, nil
}
