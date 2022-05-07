package idpconfig

import (
	"encoding/json"

	"github.com/caos/zitadel/internal/domain"
	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/repository"
)

const (
	UniqueIDPConfigNameType = "idp_config_names"
)

func NewAddIDPConfigNameUniqueConstraint(idpConfigName, resourceOwner string) *eventstore.EventUniqueConstraint {
	return eventstore.NewAddEventUniqueConstraint(
		UniqueIDPConfigNameType,
		idpConfigName+resourceOwner,
		"Errors.IDPConfig.AlreadyExists")
}

func NewRemoveIDPConfigNameUniqueConstraint(idpConfigName, resourceOwner string) *eventstore.EventUniqueConstraint {
	return eventstore.NewRemoveEventUniqueConstraint(
		UniqueIDPConfigNameType,
		idpConfigName+resourceOwner)
}

type IDPConfigAddedEvent struct {
	eventstore.BaseEvent

	ConfigID     string                      `json:"idpConfigId"`
	Name         string                      `json:"name,omitempty"`
	Typ          domain.IDPConfigType        `json:"idpType,omitempty"`
	StylingType  domain.IDPConfigStylingType `json:"stylingType,omitempty"`
	AutoRegister bool                        `json:"autoRegister,omitempty"`
}

func NewIDPConfigAddedEvent(
	base *eventstore.BaseEvent,
	configID,
	name string,
	configType domain.IDPConfigType,
	stylingType domain.IDPConfigStylingType,
	autoRegister bool,
) *IDPConfigAddedEvent {
	return &IDPConfigAddedEvent{
		BaseEvent:    *base,
		ConfigID:     configID,
		Name:         name,
		StylingType:  stylingType,
		Typ:          configType,
		AutoRegister: autoRegister,
	}
}

func (e *IDPConfigAddedEvent) Data() interface{} {
	return e
}

func (e *IDPConfigAddedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return []*eventstore.EventUniqueConstraint{NewAddIDPConfigNameUniqueConstraint(e.Name, e.Aggregate().ResourceOwner)}
}

func IDPConfigAddedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &IDPConfigAddedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "OIDC-plaBZ", "unable to unmarshal event")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}

type IDPConfigChangedEvent struct {
	eventstore.BaseEvent

	ConfigID     string                       `json:"idpConfigId"`
	Name         *string                      `json:"name,omitempty"`
	StylingType  *domain.IDPConfigStylingType `json:"stylingType,omitempty"`
	AutoRegister *bool                        `json:"autoRegister,omitempty"`
	oldName      string                       `json:"-"`
}

func (e *IDPConfigChangedEvent) Data() interface{} {
	return e
}

func (e *IDPConfigChangedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	if e.oldName == "" {
		return nil
	}
	return []*eventstore.EventUniqueConstraint{
		NewRemoveIDPConfigNameUniqueConstraint(e.oldName, e.Aggregate().ResourceOwner),
		NewAddIDPConfigNameUniqueConstraint(*e.Name, e.Aggregate().ResourceOwner),
	}
}

func NewIDPConfigChangedEvent(
	base *eventstore.BaseEvent,
	configID,
	oldName string,
	changes []IDPConfigChanges,
) (*IDPConfigChangedEvent, error) {
	if len(changes) == 0 {
		return nil, errors.ThrowPreconditionFailed(nil, "IDPCONFIG-Dsg21", "Errors.NoChangesFound")
	}
	changeEvent := &IDPConfigChangedEvent{
		BaseEvent: *base,
		ConfigID:  configID,
		oldName:   oldName,
	}
	for _, change := range changes {
		change(changeEvent)
	}
	return changeEvent, nil
}

type IDPConfigChanges func(*IDPConfigChangedEvent)

func ChangeName(name string) func(*IDPConfigChangedEvent) {
	return func(e *IDPConfigChangedEvent) {
		e.Name = &name
	}
}

func ChangeStyleType(styleType domain.IDPConfigStylingType) func(*IDPConfigChangedEvent) {
	return func(e *IDPConfigChangedEvent) {
		e.StylingType = &styleType
	}
}

func ChangeAutoRegister(autoRegister bool) func(*IDPConfigChangedEvent) {
	return func(e *IDPConfigChangedEvent) {
		e.AutoRegister = &autoRegister
	}
}

func IDPConfigChangedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &IDPConfigChangedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "OIDC-plaBZ", "unable to unmarshal event")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}

type IDPConfigDeactivatedEvent struct {
	eventstore.BaseEvent

	ConfigID string `json:"idpConfigId"`
}

func NewIDPConfigDeactivatedEvent(
	base *eventstore.BaseEvent,
	configID string,
) *IDPConfigDeactivatedEvent {

	return &IDPConfigDeactivatedEvent{
		BaseEvent: *base,
		ConfigID:  configID,
	}
}

func (e *IDPConfigDeactivatedEvent) Data() interface{} {
	return e
}

func (e *IDPConfigDeactivatedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func IDPConfigDeactivatedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &IDPConfigDeactivatedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "OIDC-plaBZ", "unable to unmarshal event")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}

type IDPConfigReactivatedEvent struct {
	eventstore.BaseEvent

	ConfigID string `json:"idpConfigId"`
}

func NewIDPConfigReactivatedEvent(
	base *eventstore.BaseEvent,
	configID string,
) *IDPConfigReactivatedEvent {

	return &IDPConfigReactivatedEvent{
		BaseEvent: *base,
		ConfigID:  configID,
	}
}

func (e *IDPConfigReactivatedEvent) Data() interface{} {
	return e
}

func (e *IDPConfigReactivatedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func IDPConfigReactivatedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &IDPConfigReactivatedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "OIDC-plaBZ", "unable to unmarshal event")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}

type IDPConfigRemovedEvent struct {
	eventstore.BaseEvent

	ConfigID string `json:"idpConfigId"`
	name     string
}

func NewIDPConfigRemovedEvent(
	base *eventstore.BaseEvent,
	configID string,
	name string,
) *IDPConfigRemovedEvent {

	return &IDPConfigRemovedEvent{
		BaseEvent: *base,
		ConfigID:  configID,
		name:      name,
	}
}

func (e *IDPConfigRemovedEvent) Data() interface{} {
	return e
}

func (e *IDPConfigRemovedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return []*eventstore.EventUniqueConstraint{NewRemoveIDPConfigNameUniqueConstraint(e.name, e.Aggregate().ResourceOwner)}
}

func IDPConfigRemovedEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &IDPConfigRemovedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "OIDC-plaBZ", "unable to unmarshal event")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}
