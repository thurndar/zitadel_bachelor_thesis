package user

import (
	"encoding/json"

	"github.com/caos/zitadel/internal/domain"
	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/repository"
)

type HumanWebAuthNAddedEvent struct {
	eventstore.BaseEvent

	WebAuthNTokenID string `json:"webAuthNTokenId"`
	Challenge       string `json:"challenge"`
}

func (e *HumanWebAuthNAddedEvent) Data() interface{} {
	return e
}

func (e *HumanWebAuthNAddedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewHumanWebAuthNAddedEvent(
	base *eventstore.BaseEvent,
	webAuthNTokenID,
	challenge string,
) *HumanWebAuthNAddedEvent {
	return &HumanWebAuthNAddedEvent{
		BaseEvent:       *base,
		WebAuthNTokenID: webAuthNTokenID,
		Challenge:       challenge,
	}
}

func HumanWebAuthNAddedEventMapper(event *repository.Event) (eventstore.Event, error) {
	webAuthNAdded := &HumanWebAuthNAddedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}
	err := json.Unmarshal(event.Data, webAuthNAdded)
	if err != nil {
		return nil, errors.ThrowInternal(err, "USER-tB8sf", "unable to unmarshal human webAuthN added")
	}
	webAuthNAdded.BaseEvent = *eventstore.BaseEventFromRepo(event)
	return webAuthNAdded, nil
}

type HumanWebAuthNVerifiedEvent struct {
	eventstore.BaseEvent

	WebAuthNTokenID   string `json:"webAuthNTokenId"`
	KeyID             []byte `json:"keyId"`
	PublicKey         []byte `json:"publicKey"`
	AttestationType   string `json:"attestationType"`
	AAGUID            []byte `json:"aaguid"`
	SignCount         uint32 `json:"signCount"`
	WebAuthNTokenName string `json:"webAuthNTokenName"`
	UserAgentID       string `json:"userAgentID,omitempty"`
}

func (e *HumanWebAuthNVerifiedEvent) Data() interface{} {
	return e
}

func (e *HumanWebAuthNVerifiedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewHumanWebAuthNVerifiedEvent(
	base *eventstore.BaseEvent,
	webAuthNTokenID,
	webAuthNTokenName,
	attestationType string,
	keyID,
	publicKey,
	aaguid []byte,
	signCount uint32,
	userAgentID string,
) *HumanWebAuthNVerifiedEvent {
	return &HumanWebAuthNVerifiedEvent{
		BaseEvent:         *base,
		WebAuthNTokenID:   webAuthNTokenID,
		KeyID:             keyID,
		PublicKey:         publicKey,
		AttestationType:   attestationType,
		AAGUID:            aaguid,
		SignCount:         signCount,
		WebAuthNTokenName: webAuthNTokenName,
		UserAgentID:       userAgentID,
	}
}

func HumanWebAuthNVerifiedEventMapper(event *repository.Event) (eventstore.Event, error) {
	webauthNVerified := &HumanWebAuthNVerifiedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}
	err := json.Unmarshal(event.Data, webauthNVerified)
	if err != nil {
		return nil, errors.ThrowInternal(err, "USER-B0zDs", "unable to unmarshal human webAuthN verified")
	}
	webauthNVerified.BaseEvent = *eventstore.BaseEventFromRepo(event)
	return webauthNVerified, nil
}

type HumanWebAuthNSignCountChangedEvent struct {
	eventstore.BaseEvent

	WebAuthNTokenID string `json:"webAuthNTokenId"`
	SignCount       uint32 `json:"signCount"`
}

func (e *HumanWebAuthNSignCountChangedEvent) Data() interface{} {
	return e
}

func (e *HumanWebAuthNSignCountChangedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewHumanWebAuthNSignCountChangedEvent(
	base *eventstore.BaseEvent,
	webAuthNTokenID string,
	signCount uint32,
) *HumanWebAuthNSignCountChangedEvent {
	return &HumanWebAuthNSignCountChangedEvent{
		BaseEvent:       *base,
		WebAuthNTokenID: webAuthNTokenID,
		SignCount:       signCount,
	}
}

func HumanWebAuthNSignCountChangedEventMapper(event *repository.Event) (eventstore.Event, error) {
	webauthNVerified := &HumanWebAuthNSignCountChangedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}
	err := json.Unmarshal(event.Data, webauthNVerified)
	if err != nil {
		return nil, errors.ThrowInternal(err, "USER-5Gm0s", "unable to unmarshal human webAuthN sign count")
	}
	webauthNVerified.BaseEvent = *eventstore.BaseEventFromRepo(event)
	return webauthNVerified, nil
}

type HumanWebAuthNRemovedEvent struct {
	eventstore.BaseEvent

	WebAuthNTokenID string          `json:"webAuthNTokenId"`
	State           domain.MFAState `json:"-"`
}

func (e *HumanWebAuthNRemovedEvent) Data() interface{} {
	return e
}

func (e *HumanWebAuthNRemovedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewHumanWebAuthNRemovedEvent(
	base *eventstore.BaseEvent,
	webAuthNTokenID string,
) *HumanWebAuthNRemovedEvent {
	return &HumanWebAuthNRemovedEvent{
		BaseEvent:       *base,
		WebAuthNTokenID: webAuthNTokenID,
	}
}

func HumanWebAuthNRemovedEventMapper(event *repository.Event) (eventstore.Event, error) {
	webauthNVerified := &HumanWebAuthNRemovedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}
	err := json.Unmarshal(event.Data, webauthNVerified)
	if err != nil {
		return nil, errors.ThrowInternal(err, "USER-gM9sd", "unable to unmarshal human webAuthN token removed")
	}
	webauthNVerified.BaseEvent = *eventstore.BaseEventFromRepo(event)
	return webauthNVerified, nil
}

type HumanWebAuthNBeginLoginEvent struct {
	eventstore.BaseEvent

	Challenge            string                             `json:"challenge"`
	AllowedCredentialIDs [][]byte                           `json:"allowedCredentialIDs"`
	UserVerification     domain.UserVerificationRequirement `json:"userVerification"`
	*AuthRequestInfo
}

func (e *HumanWebAuthNBeginLoginEvent) Data() interface{} {
	return e
}

func (e *HumanWebAuthNBeginLoginEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewHumanWebAuthNBeginLoginEvent(base *eventstore.BaseEvent, challenge string, allowedCredentialIDs [][]byte, userVerification domain.UserVerificationRequirement, info *AuthRequestInfo) *HumanWebAuthNBeginLoginEvent {
	return &HumanWebAuthNBeginLoginEvent{
		BaseEvent:            *base,
		Challenge:            challenge,
		AllowedCredentialIDs: allowedCredentialIDs,
		UserVerification:     userVerification,
		AuthRequestInfo:      info,
	}
}

func HumanWebAuthNBeginLoginEventMapper(event *repository.Event) (eventstore.Event, error) {
	webAuthNAdded := &HumanWebAuthNBeginLoginEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}
	err := json.Unmarshal(event.Data, webAuthNAdded)
	if err != nil {
		return nil, errors.ThrowInternal(err, "USER-rMb8x", "unable to unmarshal human webAuthN begin login")
	}
	webAuthNAdded.BaseEvent = *eventstore.BaseEventFromRepo(event)
	return webAuthNAdded, nil
}

type HumanWebAuthNCheckSucceededEvent struct {
	eventstore.BaseEvent
	*AuthRequestInfo
}

func (e *HumanWebAuthNCheckSucceededEvent) Data() interface{} {
	return e
}

func (e *HumanWebAuthNCheckSucceededEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewHumanWebAuthNCheckSucceededEvent(
	base *eventstore.BaseEvent,
	info *AuthRequestInfo) *HumanWebAuthNCheckSucceededEvent {
	return &HumanWebAuthNCheckSucceededEvent{
		BaseEvent:       *base,
		AuthRequestInfo: info,
	}
}

func HumanWebAuthNCheckSucceededEventMapper(event *repository.Event) (eventstore.Event, error) {
	webAuthNAdded := &HumanWebAuthNCheckSucceededEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}
	err := json.Unmarshal(event.Data, webAuthNAdded)
	if err != nil {
		return nil, errors.ThrowInternal(err, "USER-2M0fg", "unable to unmarshal human webAuthN check succeeded")
	}
	webAuthNAdded.BaseEvent = *eventstore.BaseEventFromRepo(event)
	return webAuthNAdded, nil
}

type HumanWebAuthNCheckFailedEvent struct {
	eventstore.BaseEvent
	*AuthRequestInfo
}

func (e *HumanWebAuthNCheckFailedEvent) Data() interface{} {
	return e
}

func (e *HumanWebAuthNCheckFailedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewHumanWebAuthNCheckFailedEvent(
	base *eventstore.BaseEvent,
	info *AuthRequestInfo) *HumanWebAuthNCheckFailedEvent {
	return &HumanWebAuthNCheckFailedEvent{
		BaseEvent:       *base,
		AuthRequestInfo: info,
	}
}

func HumanWebAuthNCheckFailedEventMapper(event *repository.Event) (eventstore.Event, error) {
	webAuthNAdded := &HumanWebAuthNCheckFailedEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}
	err := json.Unmarshal(event.Data, webAuthNAdded)
	if err != nil {
		return nil, errors.ThrowInternal(err, "USER-O0dse", "unable to unmarshal human webAuthN check failed")
	}
	webAuthNAdded.BaseEvent = *eventstore.BaseEventFromRepo(event)
	return webAuthNAdded, nil
}
