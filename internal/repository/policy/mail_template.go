package policy

import (
	"encoding/json"

	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/repository"
)

const (
	mailPolicyPrefix                   = "mail."
	mailTemplatePolicyPrefix           = mailPolicyPrefix + "template."
	MailTemplatePolicyAddedEventType   = mailTemplatePolicyPrefix + "added"
	MailTemplatePolicyChangedEventType = mailTemplatePolicyPrefix + "changed"
	MailTemplatePolicyRemovedEventType = mailTemplatePolicyPrefix + "removed"
)

type MailTemplateAddedEvent struct {
	eventstore.BaseEvent

	Template []byte `json:"template,omitempty"`
}

func (e *MailTemplateAddedEvent) Data() interface{} {
	return e
}

func (e *MailTemplateAddedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewMailTemplateAddedEvent(
	base *eventstore.BaseEvent,
	template []byte,
) *MailTemplateAddedEvent {
	return &MailTemplateAddedEvent{
		BaseEvent: *base,
		Template:  template,
	}
}

func MailTemplateAddedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &MailTemplateAddedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "POLIC-5m9if", "unable to unmarshal mail template")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}

type MailTemplateChangedEvent struct {
	eventstore.BaseEvent

	Template *[]byte `json:"template,omitempty"`
}

func (e *MailTemplateChangedEvent) Data() interface{} {
	return e
}

func (e *MailTemplateChangedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewMailTemplateChangedEvent(
	base *eventstore.BaseEvent,
	changes []MailTemplateChanges,
) (*MailTemplateChangedEvent, error) {
	if len(changes) == 0 {
		return nil, errors.ThrowPreconditionFailed(nil, "POLICY-m9osd", "Errors.NoChangesFound")
	}
	changeEvent := &MailTemplateChangedEvent{
		BaseEvent: *base,
	}
	for _, change := range changes {
		change(changeEvent)
	}
	return changeEvent, nil
}

type MailTemplateChanges func(*MailTemplateChangedEvent)

func ChangeTemplate(template []byte) func(*MailTemplateChangedEvent) {
	return func(e *MailTemplateChangedEvent) {
		e.Template = &template
	}
}

func MailTemplateChangedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &MailTemplateChangedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "POLIC-3uu8K", "unable to unmarshal mail template policy")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}

type MailTemplateRemovedEvent struct {
	eventstore.BaseEvent
}

func (e *MailTemplateRemovedEvent) Data() interface{} {
	return nil
}

func (e *MailTemplateRemovedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewMailTemplateRemovedEvent(base *eventstore.BaseEvent) *MailTemplateRemovedEvent {
	return &MailTemplateRemovedEvent{
		BaseEvent: *base,
	}
}

func MailTemplateRemovedEventMapper(event *repository.Event) (eventstore.Event, error) {
	return &MailTemplateRemovedEvent{
		BaseEvent: *eventstore.BaseEventFromRepo(event),
	}, nil
}
