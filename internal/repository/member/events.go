package member

import (
	"encoding/json"
	"fmt"

	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/repository"
)

const (
	UniqueMember            = "member"
	AddedEventType          = "member.added"
	ChangedEventType        = "member.changed"
	RemovedEventType        = "member.removed"
	CascadeRemovedEventType = "member.cascade.removed"
)

func NewAddMemberUniqueConstraint(aggregateID, userID string) *eventstore.EventUniqueConstraint {
	return eventstore.NewAddEventUniqueConstraint(
		UniqueMember,
		fmt.Sprintf("%s:%s", aggregateID, userID),
		"Errors.Member.AlreadyExists")
}

func NewRemoveMemberUniqueConstraint(aggregateID, userID string) *eventstore.EventUniqueConstraint {
	return eventstore.NewRemoveEventUniqueConstraint(
		UniqueMember,
		fmt.Sprintf("%s:%s", aggregateID, userID),
	)
}

type MemberAddedEvent struct {
	eventstore.BaseEvent

	Roles  []string `json:"roles"`
	UserID string   `json:"userId"`
}

func (e *MemberAddedEvent) Data() interface{} {
	return e
}

func (e *MemberAddedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return []*eventstore.EventUniqueConstraint{NewAddMemberUniqueConstraint(e.Aggregate().ID, e.UserID)}
}

func NewMemberAddedEvent(
	base *eventstore.BaseEvent,
	userID string,
	roles ...string,
) *MemberAddedEvent {

	return &MemberAddedEvent{
		BaseEvent: *base,
		Roles:     roles,
		UserID:    userID,
	}
}

func MemberAddedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &MemberAddedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "POLIC-puqv4", "unable to unmarshal label policy")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}

type MemberChangedEvent struct {
	eventstore.BaseEvent

	Roles  []string `json:"roles,omitempty"`
	UserID string   `json:"userId,omitempty"`
}

func (e *MemberChangedEvent) Data() interface{} {
	return e
}

func (e *MemberChangedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewMemberChangedEvent(
	base *eventstore.BaseEvent,
	userID string,
	roles ...string,
) *MemberChangedEvent {
	return &MemberChangedEvent{
		BaseEvent: *base,
		Roles:     roles,
		UserID:    userID,
	}
}

func ChangedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &MemberChangedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "POLIC-puqv4", "unable to unmarshal label policy")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}

type MemberRemovedEvent struct {
	eventstore.BaseEvent

	UserID string `json:"userId"`
}

func (e *MemberRemovedEvent) Data() interface{} {
	return e
}

func (e *MemberRemovedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return []*eventstore.EventUniqueConstraint{NewRemoveMemberUniqueConstraint(e.Aggregate().ID, e.UserID)}
}

func NewRemovedEvent(
	base *eventstore.BaseEvent,
	userID string,
) *MemberRemovedEvent {

	return &MemberRemovedEvent{
		BaseEvent: *base,
		UserID:    userID,
	}
}

func RemovedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &MemberRemovedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "MEMBER-Ep4ip", "unable to unmarshal label policy")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}

type MemberCascadeRemovedEvent struct {
	eventstore.BaseEvent

	UserID string `json:"userId"`
}

func (e *MemberCascadeRemovedEvent) Data() interface{} {
	return e
}

func (e *MemberCascadeRemovedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return []*eventstore.EventUniqueConstraint{NewRemoveMemberUniqueConstraint(e.Aggregate().ID, e.UserID)}
}

func NewCascadeRemovedEvent(
	base *eventstore.BaseEvent,
	userID string,
) *MemberCascadeRemovedEvent {

	return &MemberCascadeRemovedEvent{
		BaseEvent: *base,
		UserID:    userID,
	}
}

func CascadeRemovedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &MemberCascadeRemovedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "MEMBER-3j9sf", "unable to unmarshal label policy")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}
