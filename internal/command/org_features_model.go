package command

import (
	"context"
	"time"

	"github.com/caos/zitadel/internal/domain"
	"github.com/caos/zitadel/internal/eventstore"
	"github.com/caos/zitadel/internal/repository/features"
	"github.com/caos/zitadel/internal/repository/org"
)

type OrgFeaturesWriteModel struct {
	FeaturesWriteModel
}

func NewOrgFeaturesWriteModel(orgID string) *OrgFeaturesWriteModel {
	return &OrgFeaturesWriteModel{
		FeaturesWriteModel{
			WriteModel: eventstore.WriteModel{
				AggregateID:   orgID,
				ResourceOwner: orgID,
			},
		},
	}
}

func (wm *OrgFeaturesWriteModel) AppendEvents(events ...eventstore.Event) {
	for _, event := range events {
		switch e := event.(type) {
		case *org.FeaturesSetEvent:
			wm.FeaturesWriteModel.AppendEvents(&e.FeaturesSetEvent)
		case *org.FeaturesRemovedEvent:
			wm.FeaturesWriteModel.AppendEvents(&e.FeaturesRemovedEvent)
		}
	}
}

func (wm *OrgFeaturesWriteModel) IsValid() bool {
	return wm.AggregateID != ""
}

func (wm *OrgFeaturesWriteModel) Reduce() error {
	return wm.FeaturesWriteModel.Reduce()
}

func (wm *OrgFeaturesWriteModel) Query() *eventstore.SearchQueryBuilder {
	return eventstore.NewSearchQueryBuilder(eventstore.ColumnsEvent).
		ResourceOwner(wm.ResourceOwner).
		AddQuery().
		AggregateTypes(org.AggregateType).
		AggregateIDs(wm.FeaturesWriteModel.AggregateID).
		EventTypes(
			org.FeaturesSetEventType,
			org.FeaturesRemovedEventType).
		Builder()
}

func (wm *OrgFeaturesWriteModel) NewSetEvent(
	ctx context.Context,
	aggregate *eventstore.Aggregate,
	tierName,
	tierDescription string,
	state domain.FeaturesState,
	stateDescription string,
	auditLogRetention time.Duration,
	loginPolicyFactors,
	loginPolicyIDP,
	loginPolicyPasswordless,
	loginPolicyRegistration,
	loginPolicyUsernameLogin,
	loginPolicyPasswordReset,
	passwordComplexityPolicy,
	labelPolicyPrivateLabel,
	labelPolicyWatermark,
	customDomain,
	privacyPolicy,
	metadataUser,
	customTextMessage,
	customTextLogin,
	lockoutPolicy bool,
	actionsAllowed domain.ActionsAllowed,
	maxActions int,
) (*org.FeaturesSetEvent, bool) {

	changes := make([]features.FeaturesChanges, 0)

	if tierName != "" && wm.TierName != tierName {
		changes = append(changes, features.ChangeTierName(tierName))
	}
	if wm.TierDescription != tierDescription {
		changes = append(changes, features.ChangeTierDescription(tierDescription))
	}
	if wm.State != state {
		changes = append(changes, features.ChangeState(state))
	}
	if stateDescription != "" && wm.StateDescription != stateDescription {
		changes = append(changes, features.ChangeStateDescription(stateDescription))
	}
	if auditLogRetention != 0 && wm.AuditLogRetention != auditLogRetention {
		changes = append(changes, features.ChangeAuditLogRetention(auditLogRetention))
	}
	if wm.LoginPolicyFactors != loginPolicyFactors {
		changes = append(changes, features.ChangeLoginPolicyFactors(loginPolicyFactors))
	}
	if wm.LoginPolicyIDP != loginPolicyIDP {
		changes = append(changes, features.ChangeLoginPolicyIDP(loginPolicyIDP))
	}
	if wm.LoginPolicyPasswordless != loginPolicyPasswordless {
		changes = append(changes, features.ChangeLoginPolicyPasswordless(loginPolicyPasswordless))
	}
	if wm.LoginPolicyRegistration != loginPolicyRegistration {
		changes = append(changes, features.ChangeLoginPolicyRegistration(loginPolicyRegistration))
	}
	if wm.LoginPolicyUsernameLogin != loginPolicyUsernameLogin {
		changes = append(changes, features.ChangeLoginPolicyUsernameLogin(loginPolicyUsernameLogin))
	}
	if wm.LoginPolicyPasswordReset != loginPolicyPasswordReset {
		changes = append(changes, features.ChangeLoginPolicyPasswordReset(loginPolicyPasswordReset))
	}
	if wm.PasswordComplexityPolicy != passwordComplexityPolicy {
		changes = append(changes, features.ChangePasswordComplexityPolicy(passwordComplexityPolicy))
	}
	if wm.LabelPolicyPrivateLabel != labelPolicyPrivateLabel {
		changes = append(changes, features.ChangeLabelPolicyPrivateLabel(labelPolicyPrivateLabel))
	}
	if wm.LabelPolicyWatermark != labelPolicyWatermark {
		changes = append(changes, features.ChangeLabelPolicyWatermark(labelPolicyWatermark))
	}
	if wm.CustomDomain != customDomain {
		changes = append(changes, features.ChangeCustomDomain(customDomain))
	}
	if wm.PrivacyPolicy != privacyPolicy {
		changes = append(changes, features.ChangePrivacyPolicy(privacyPolicy))
	}
	if wm.MetadataUser != metadataUser {
		changes = append(changes, features.ChangeMetadataUser(metadataUser))
	}
	if wm.CustomTextMessage != customTextMessage {
		changes = append(changes, features.ChangeCustomTextMessage(customTextMessage))
	}
	if wm.CustomTextLogin != customTextLogin {
		changes = append(changes, features.ChangeCustomTextLogin(customTextLogin))
	}
	if wm.LockoutPolicy != lockoutPolicy {
		changes = append(changes, features.ChangeLockoutPolicy(lockoutPolicy))
	}
	if wm.ActionsAllowed != actionsAllowed {
		changes = append(changes, features.ChangeActionsAllowed(actionsAllowed))
	}
	if wm.MaxActions != maxActions {
		changes = append(changes, features.ChangeMaxActions(maxActions))
	}

	if len(changes) == 0 {
		return nil, false
	}
	changedEvent, err := org.NewFeaturesSetEvent(ctx, aggregate, changes)
	if err != nil {
		return nil, false
	}
	return changedEvent, true
}
