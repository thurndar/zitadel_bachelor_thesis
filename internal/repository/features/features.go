package features

import (
	"encoding/json"
	"time"

	"github.com/caos/zitadel/internal/domain"
	"github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/eventstore/repository"
)

const (
	featuresPrefix           = "features."
	FeaturesSetEventType     = featuresPrefix + "set"
	FeaturesRemovedEventType = featuresPrefix + "removed"
)

type FeaturesSetEvent struct {
	eventstore.BaseEvent

	TierName                 *string                `json:"tierName,omitempty"`
	TierDescription          *string                `json:"tierDescription,omitempty"`
	State                    *domain.FeaturesState  `json:"state,omitempty"`
	StateDescription         *string                `json:"stateDescription,omitempty"`
	AuditLogRetention        *time.Duration         `json:"auditLogRetention,omitempty"`
	LoginPolicyFactors       *bool                  `json:"loginPolicyFactors,omitempty"`
	LoginPolicyIDP           *bool                  `json:"loginPolicyIDP,omitempty"`
	LoginPolicyPasswordless  *bool                  `json:"loginPolicyPasswordless,omitempty"`
	LoginPolicyRegistration  *bool                  `json:"loginPolicyRegistration,omitempty"`
	LoginPolicyUsernameLogin *bool                  `json:"loginPolicyUsernameLogin,omitempty"`
	LoginPolicyPasswordReset *bool                  `json:"loginPolicyPasswordReset,omitempty"`
	PasswordComplexityPolicy *bool                  `json:"passwordComplexityPolicy,omitempty"`
	LabelPolicy              *bool                  `json:"labelPolicy,omitempty"`
	LabelPolicyPrivateLabel  *bool                  `json:"labelPolicyPrivateLabel,omitempty"`
	LabelPolicyWatermark     *bool                  `json:"labelPolicyWatermark,omitempty"`
	CustomDomain             *bool                  `json:"customDomain,omitempty"`
	PrivacyPolicy            *bool                  `json:"privacyPolicy,omitempty"`
	MetadataUser             *bool                  `json:"metadataUser,omitempty"`
	CustomTextMessage        *bool                  `json:"customTextMessage,omitempty"`
	CustomTextLogin          *bool                  `json:"customTextLogin,omitempty"`
	LockoutPolicy            *bool                  `json:"lockoutPolicy,omitempty"`
	Actions                  *bool                  `json:"actions,omitempty"`
	ActionsAllowed           *domain.ActionsAllowed `json:"actionsAllowed,omitempty"`
	MaxActions               *int                   `json:"maxActions,omitempty"`
}

func (e *FeaturesSetEvent) Data() interface{} {
	return e
}

func (e *FeaturesSetEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func NewFeaturesSetEvent(
	base *eventstore.BaseEvent,
	changes []FeaturesChanges,
) (*FeaturesSetEvent, error) {
	if len(changes) == 0 {
		return nil, errors.ThrowPreconditionFailed(nil, "FEATURES-d34F4", "Errors.NoChangesFound")
	}
	changeEvent := &FeaturesSetEvent{
		BaseEvent: *base,
	}
	for _, change := range changes {
		change(changeEvent)
	}
	return changeEvent, nil
}

type FeaturesChanges func(*FeaturesSetEvent)

func ChangeTierName(tierName string) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.TierName = &tierName
	}
}

func ChangeTierDescription(tierDescription string) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.TierDescription = &tierDescription
	}
}

func ChangeState(State domain.FeaturesState) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.State = &State
	}
}

func ChangeStateDescription(statusDescription string) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.StateDescription = &statusDescription
	}
}

func ChangeAuditLogRetention(retention time.Duration) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.AuditLogRetention = &retention
	}
}

func ChangeLoginPolicyFactors(loginPolicyFactors bool) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.LoginPolicyFactors = &loginPolicyFactors
	}
}

func ChangeLoginPolicyIDP(loginPolicyIDP bool) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.LoginPolicyIDP = &loginPolicyIDP
	}
}

func ChangeLoginPolicyPasswordless(loginPolicyPasswordless bool) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.LoginPolicyPasswordless = &loginPolicyPasswordless
	}
}

func ChangeLoginPolicyRegistration(loginPolicyRegistration bool) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.LoginPolicyRegistration = &loginPolicyRegistration
	}
}

func ChangeLoginPolicyUsernameLogin(loginPolicyUsernameLogin bool) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.LoginPolicyUsernameLogin = &loginPolicyUsernameLogin
	}
}

func ChangeLoginPolicyPasswordReset(loginPolicyPasswordReset bool) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.LoginPolicyPasswordReset = &loginPolicyPasswordReset
	}
}

func ChangePasswordComplexityPolicy(passwordComplexityPolicy bool) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.PasswordComplexityPolicy = &passwordComplexityPolicy
	}
}

func ChangeLabelPolicyPrivateLabel(labelPolicyPrivateLabel bool) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.LabelPolicyPrivateLabel = &labelPolicyPrivateLabel
	}
}

func ChangeLabelPolicyWatermark(labelPolicyWatermark bool) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.LabelPolicyWatermark = &labelPolicyWatermark
	}
}

func ChangeCustomDomain(customDomain bool) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.CustomDomain = &customDomain
	}
}

func ChangePrivacyPolicy(privacyPolicy bool) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.PrivacyPolicy = &privacyPolicy
	}
}

func ChangeMetadataUser(metadataUser bool) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.MetadataUser = &metadataUser
	}
}

func ChangeCustomTextMessage(customTextMessage bool) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.CustomTextMessage = &customTextMessage
	}
}

func ChangeCustomTextLogin(customTextLogin bool) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.CustomTextLogin = &customTextLogin
	}
}

func ChangeLockoutPolicy(lockoutPolicy bool) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.LockoutPolicy = &lockoutPolicy
	}
}

func ChangeActionsAllowed(allowedType domain.ActionsAllowed) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.ActionsAllowed = &allowedType
	}
}

func ChangeMaxActions(maxActions int) func(event *FeaturesSetEvent) {
	return func(e *FeaturesSetEvent) {
		e.MaxActions = &maxActions
	}
}

func FeaturesSetEventMapper(event *repository.Event) (eventstore.Event, error) {
	e := &FeaturesSetEvent{
		// BaseEvent: *eventstore.BaseEventFromRepo(event),
	}

	err := json.Unmarshal(event.Data, e)
	if err != nil {
		return nil, errors.ThrowInternal(err, "FEATURES-fdgDg", "unable to unmarshal features")
	}
	e.BaseEvent = *eventstore.BaseEventFromRepo(event)

	return e, nil
}

type FeaturesRemovedEvent struct {
	eventstore.BaseEvent
}

func (e *FeaturesRemovedEvent) Data() interface{} {
	return nil
}

func (e *FeaturesRemovedEvent) UniqueConstraints() []*eventstore.EventUniqueConstraint {
	return nil
}

func (e *FeaturesRemovedEvent) Assets() []*eventstore.Asset {
	return nil
}

func NewFeaturesRemovedEvent(base *eventstore.BaseEvent) *FeaturesRemovedEvent {
	return &FeaturesRemovedEvent{
		BaseEvent: *base,
	}
}

func FeaturesRemovedEventMapper(event *repository.Event) (eventstore.Event, error) {
	return &FeaturesRemovedEvent{
		BaseEvent: *eventstore.BaseEventFromRepo(event),
	}, nil
}
