package project

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/caos/zitadel/internal/eventstore"

	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore/repository"
)

var (
	UniqueRoleType      = "project_role"
	roleEventTypePrefix = projectEventTypePrefix + "role."
	RoleAddedType       = roleEventTypePrefix + "added"
	RoleChangedType     = roleEventTypePrefix + "changed"
	RoleRemovedType     = roleEventTypePrefix + "removed"
)

func NewAddProjectRoleUniqueConstraint(roleKey, projectID string) *eventstore.EventUniqueConstraint {
	return eventstore.NewAddEventUniqueConstraint(
		UniqueRoleType,
		fmt.Sprintf("%s:%s", roleKey, projectID),
		"Errors.Project.Role.AlreadyExists")
}

func NewRemoveProjectRoleUniqueConstraint(roleKey, projectID string) *eventstore.EventUniqueConstraint {
	return eventstore.NewRemoveEventUniqueConstraint(
		UniqueRoleType,
		fmt.Sprintf("%s:%s", roleKey, projectID))
}

type RoleAddedEvent struct {
	eventstore.BaseEvent

	Key         string `json:"key,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Group       string `json:"group,omitempty"`
}

func (e *RoleAddedEvent) Data() interface{} {
	return e
}

func (e *RoleAddedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return []*eventstore.EventUniqueConstraint{NewAddProjectRoleUniqueConstraint(e.Key, e.Aggregate().ID)}
}

func NewRoleAddedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	key,
	displayName,
	group string,
) *RoleAddedEvent {
	return &RoleAddedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			RoleAddedType,
		),
		Key:         key,
		DisplayName: displayName,
		Group:       group,
	}
}

func RoleAddedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &RoleAddedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "PROJECT-2M0xy", "unable to unmarshal project role")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}

type RoleChangedEvent struct {
	eventstore.BaseEvent

	Key         string  `json:"key,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Group       *string `json:"group,omitempty"`
}

func (e *RoleChangedEvent) Data() interface{} {
	return e
}

func (e *RoleChangedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewRoleChangedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	key string,
	changes []RoleChanges,
) (*RoleChangedEvent, error) {
	if len(changes) == 0 {
		return nil, errors.ThrowPreconditionFailed(nil, "PROJECT-eR9vx", "Errors.NoChangesFound")
	}
	changeEvent := &RoleChangedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			RoleChangedType,
		),
		Key: key,
	}
	for _, change := range changes {
		change(changeEvent)
	}
	return changeEvent, nil
}

type RoleChanges func(event *RoleChangedEvent)

func ChangeKey(key string) func(event *RoleChangedEvent) {
	return func(e *RoleChangedEvent) {
		e.Key = key
	}
}

func ChangeDisplayName(displayName string) func(event *RoleChangedEvent) {
	return func(e *RoleChangedEvent) {
		e.DisplayName = &displayName
	}
}

func ChangeGroup(group string) func(event *RoleChangedEvent) {
	return func(e *RoleChangedEvent) {
		e.Group = &group
	}
}
func RoleChangedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &RoleChangedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "PROJECT-3M0vx", "unable to unmarshal project role")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}

type RoleRemovedEvent struct {
	eventstore.BaseEvent

	Key string `json:"key,omitempty"`
}

func (e *RoleRemovedEvent) Data() interface{} {
	return e
}

func (e *RoleRemovedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return []*eventstore.EventUniqueConstraint{NewRemoveProjectRoleUniqueConstraint(e.Key, e.Aggregate().ID)}
}

func NewRoleRemovedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	key string) *RoleRemovedEvent {
	return &RoleRemovedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			RoleRemovedType,
		),
		Key: key,
	}
}

func RoleRemovedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &RoleRemovedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "PROJECT-1M0xs", "unable to unmarshal project role")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}
