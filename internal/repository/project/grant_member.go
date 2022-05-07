package project

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/caos/zitadel/internal/eventstore"

	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore/repository"
	"github.com/caos/zitadel/internal/repository/member"
)

var (
	UniqueProjectGrantMemberType  = "project_grant_member"
	GrantMemberAddedType          = grantEventTypePrefix + member.AddedEventType
	GrantMemberChangedType        = grantEventTypePrefix + member.ChangedEventType
	GrantMemberRemovedType        = grantEventTypePrefix + member.RemovedEventType
	GrantMemberCascadeRemovedType = grantEventTypePrefix + member.CascadeRemovedEventType
)

func NewAddProjectGrantMemberUniqueConstraint(projectID, userID, grantID string) *eventstore.EventUniqueConstraint {
	return eventstore.NewAddEventUniqueConstraint(
		UniqueProjectGrantMemberType,
		fmt.Sprintf("%s:%s:%s", projectID, userID, grantID),
		"Errors.Project.Member.AlreadyExists")
}

func NewRemoveProjectGrantMemberUniqueConstraint(projectID, userID, grantID string) *eventstore.EventUniqueConstraint {
	return eventstore.NewRemoveEventUniqueConstraint(
		UniqueProjectGrantMemberType,
		fmt.Sprintf("%s:%s:%s", projectID, userID, grantID),
	)
}

type GrantMemberAddedEvent struct {
	eventstore.BaseEvent

	Roles   []string `json:"roles"`
	UserID  string   `json:"userId"`
	GrantID string   `json:"grantId"`
}

func (e *GrantMemberAddedEvent) Data() interface{} {
	return e
}

func (e *GrantMemberAddedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return []*eventstore.EventUniqueConstraint{NewAddProjectGrantMemberUniqueConstraint(e.Aggregate().ID, e.UserID, e.GrantID)}
}

func NewProjectGrantMemberAddedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	userID,
	grantID string,
	roles ...string,
) *GrantMemberAddedEvent {
	return &GrantMemberAddedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			GrantMemberAddedType,
		),
		UserID:  userID,
		GrantID: grantID,
		Roles:   roles,
	}
}

func GrantMemberAddedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &GrantMemberAddedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "PROJECT-9f0sf", "unable to unmarshal label policy")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}

type GrantMemberChangedEvent struct {
	eventstore.BaseEvent

	Roles   []string `json:"roles"`
	GrantID string   `json:"grantId"`
	UserID  string   `json:"userId"`
}

func (e *GrantMemberChangedEvent) Data() interface{} {
	return e
}

func (e *GrantMemberChangedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewProjectGrantMemberChangedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	userID,
	grantID string,
	roles ...string,
) *GrantMemberChangedEvent {
	return &GrantMemberChangedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			GrantMemberAddedType,
		),
		UserID:  userID,
		GrantID: grantID,
		Roles:   roles,
	}
}

func GrantMemberChangedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &GrantMemberChangedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "PROJECT-39fi8", "unable to unmarshal label policy")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}

type GrantMemberRemovedEvent struct {
	eventstore.BaseEvent

	UserID  string `json:"userId"`
	GrantID string `json:"grantId"`
}

func (e *GrantMemberRemovedEvent) Data() interface{} {
	return e
}

func (e *GrantMemberRemovedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return []*eventstore.EventUniqueConstraint{NewRemoveProjectGrantMemberUniqueConstraint(e.Aggregate().ID, e.UserID, e.GrantID)}
}

func NewProjectGrantMemberRemovedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	userID,
	grantID string,
) *GrantMemberRemovedEvent {
	return &GrantMemberRemovedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			GrantMemberRemovedType,
		),
		UserID:  userID,
		GrantID: grantID,
	}
}

func GrantMemberRemovedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &GrantMemberRemovedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "PROJECT-173fM", "unable to unmarshal label policy")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}

type GrantMemberCascadeRemovedEvent struct {
	eventstore.BaseEvent

	UserID  string `json:"userId"`
	GrantID string `json:"grantId"`
}

func (e *GrantMemberCascadeRemovedEvent) Data() interface{} {
	return e
}

func (e *GrantMemberCascadeRemovedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return []*eventstore.EventUniqueConstraint{NewRemoveProjectGrantMemberUniqueConstraint(e.Aggregate().ID, e.UserID, e.GrantID)}
}

func NewProjectGrantMemberCascadeRemovedEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	userID,
	grantID string,
) *GrantMemberCascadeRemovedEvent {
	return &GrantMemberCascadeRemovedEvent{
		BaseEvent: *eventstore.NewBaseEventForPush(
			ctx,
			aggregate,
			GrantMemberCascadeRemovedType,
		),
		UserID:  userID,
		GrantID: grantID,
	}
}

func GrantMemberCascadeRemovedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &GrantMemberCascadeRemovedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "PROJECT-3kfs3", "unable to unmarshal label policy")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}
