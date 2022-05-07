package iam

import (
	"github.com/caos/zitadel/internal/eventstore"
)

func RegisterEventMappers(es *eventstore.Eventstore) {
	es.RegisterFilterEventMapper(SetupStartedEventType, SetupStepMapper).
		RegisterFilterEventMapper(SetupDoneEventType, SetupStepMapper).
		RegisterFilterEventMapper(GlobalOrgSetEventType, GlobalOrgSetMapper).
		RegisterFilterEventMapper(ProjectSetEventType, ProjectSetMapper).
		RegisterFilterEventMapper(UniqueConstraintsMigratedEventType, MigrateUniqueConstraintEventMapper).
		RegisterFilterEventMapper(LabelPolicyAddedEventType, LabelPolicyAddedEventMapper).
		RegisterFilterEventMapper(LabelPolicyChangedEventType, LabelPolicyChangedEventMapper).
		RegisterFilterEventMapper(LabelPolicyActivatedEventType, LabelPolicyActivatedEventMapper).
		RegisterFilterEventMapper(LabelPolicyLogoAddedEventType, LabelPolicyLogoAddedEventMapper).
		RegisterFilterEventMapper(LabelPolicyLogoRemovedEventType, LabelPolicyLogoRemovedEventMapper).
		RegisterFilterEventMapper(LabelPolicyIconAddedEventType, LabelPolicyIconAddedEventMapper).
		RegisterFilterEventMapper(LabelPolicyIconRemovedEventType, LabelPolicyIconRemovedEventMapper).
		RegisterFilterEventMapper(LabelPolicyLogoDarkAddedEventType, LabelPolicyLogoDarkAddedEventMapper).
		RegisterFilterEventMapper(LabelPolicyLogoDarkRemovedEventType, LabelPolicyLogoDarkRemovedEventMapper).
		RegisterFilterEventMapper(LabelPolicyIconDarkAddedEventType, LabelPolicyIconDarkAddedEventMapper).
		RegisterFilterEventMapper(LabelPolicyIconDarkRemovedEventType, LabelPolicyIconDarkRemovedEventMapper).
		RegisterFilterEventMapper(LabelPolicyFontAddedEventType, LabelPolicyFontAddedEventMapper).
		RegisterFilterEventMapper(LabelPolicyFontRemovedEventType, LabelPolicyFontRemovedEventMapper).
		RegisterFilterEventMapper(LabelPolicyAssetsRemovedEventType, LabelPolicyAssetsRemovedEventMapper).
		RegisterFilterEventMapper(LoginPolicyAddedEventType, LoginPolicyAddedEventMapper).
		RegisterFilterEventMapper(LoginPolicyChangedEventType, LoginPolicyChangedEventMapper).
		RegisterFilterEventMapper(OrgIAMPolicyAddedEventType, OrgIAMPolicyAddedEventMapper).
		RegisterFilterEventMapper(OrgIAMPolicyChangedEventType, OrgIAMPolicyChangedEventMapper).
		RegisterFilterEventMapper(PasswordAgePolicyAddedEventType, PasswordAgePolicyAddedEventMapper).
		RegisterFilterEventMapper(PasswordAgePolicyChangedEventType, PasswordAgePolicyChangedEventMapper).
		RegisterFilterEventMapper(PasswordComplexityPolicyAddedEventType, PasswordComplexityPolicyAddedEventMapper).
		RegisterFilterEventMapper(PasswordComplexityPolicyChangedEventType, PasswordComplexityPolicyChangedEventMapper).
		RegisterFilterEventMapper(LockoutPolicyAddedEventType, LockoutPolicyAddedEventMapper).
		RegisterFilterEventMapper(LockoutPolicyChangedEventType, LockoutPolicyChangedEventMapper).
		RegisterFilterEventMapper(PrivacyPolicyAddedEventType, PrivacyPolicyAddedEventMapper).
		RegisterFilterEventMapper(PrivacyPolicyChangedEventType, PrivacyPolicyChangedEventMapper).
		RegisterFilterEventMapper(MemberAddedEventType, MemberAddedEventMapper).
		RegisterFilterEventMapper(MemberChangedEventType, MemberChangedEventMapper).
		RegisterFilterEventMapper(MemberRemovedEventType, MemberRemovedEventMapper).
		RegisterFilterEventMapper(MemberCascadeRemovedEventType, MemberCascadeRemovedEventMapper).
		RegisterFilterEventMapper(IDPConfigAddedEventType, IDPConfigAddedEventMapper).
		RegisterFilterEventMapper(IDPConfigChangedEventType, IDPConfigChangedEventMapper).
		RegisterFilterEventMapper(IDPConfigRemovedEventType, IDPConfigRemovedEventMapper).
		RegisterFilterEventMapper(IDPConfigDeactivatedEventType, IDPConfigDeactivatedEventMapper).
		RegisterFilterEventMapper(IDPConfigReactivatedEventType, IDPConfigReactivatedEventMapper).
		RegisterFilterEventMapper(IDPOIDCConfigAddedEventType, IDPOIDCConfigAddedEventMapper).
		RegisterFilterEventMapper(IDPOIDCConfigChangedEventType, IDPOIDCConfigChangedEventMapper).
		RegisterFilterEventMapper(IDPJWTConfigAddedEventType, IDPJWTConfigAddedEventMapper).
		RegisterFilterEventMapper(IDPJWTConfigChangedEventType, IDPJWTConfigChangedEventMapper).
		RegisterFilterEventMapper(LoginPolicyIDPProviderAddedEventType, IdentityProviderAddedEventMapper).
		RegisterFilterEventMapper(LoginPolicyIDPProviderRemovedEventType, IdentityProviderRemovedEventMapper).
		RegisterFilterEventMapper(LoginPolicyIDPProviderCascadeRemovedEventType, IdentityProviderCascadeRemovedEventMapper).
		RegisterFilterEventMapper(LoginPolicySecondFactorAddedEventType, SecondFactorAddedEventMapper).
		RegisterFilterEventMapper(LoginPolicySecondFactorRemovedEventType, SecondFactorRemovedEventMapper).
		RegisterFilterEventMapper(LoginPolicyMultiFactorAddedEventType, MultiFactorAddedEventMapper).
		RegisterFilterEventMapper(LoginPolicyMultiFactorRemovedEventType, MultiFactorRemovedEventMapper).
		RegisterFilterEventMapper(MailTemplateAddedEventType, MailTemplateAddedEventMapper).
		RegisterFilterEventMapper(MailTemplateChangedEventType, MailTemplateChangedEventMapper).
		RegisterFilterEventMapper(MailTextAddedEventType, MailTextAddedEventMapper).
		RegisterFilterEventMapper(MailTextChangedEventType, MailTextChangedEventMapper).
		RegisterFilterEventMapper(CustomTextSetEventType, CustomTextSetEventMapper).
		RegisterFilterEventMapper(CustomTextRemovedEventType, CustomTextRemovedEventMapper).
		RegisterFilterEventMapper(CustomTextTemplateRemovedEventType, CustomTextTemplateRemovedEventMapper).
		RegisterFilterEventMapper(FeaturesSetEventType, FeaturesSetEventMapper)

	eventstore.RegisterEventType(SetupStartedEventType, func() eventstore.Event { return &SetupStepEvent{Done: false} })
	eventstore.RegisterEventType(SetupDoneEventType, func() eventstore.Event { return &SetupStepEvent{Done: true} })
	eventstore.RegisterEventType(GlobalOrgSetEventType, func() eventstore.Event { return new(GlobalOrgSetEvent) })
	eventstore.RegisterEventType(ProjectSetEventType, func() eventstore.Event { return new(ProjectSetEvent) })
	eventstore.RegisterEventType(UniqueConstraintsMigratedEventType, func() eventstore.Event { return new(MigrateUniqueConstraintEvent) })
	eventstore.RegisterEventType(LabelPolicyAddedEventType, func() eventstore.Event { return new(LabelPolicyAddedEvent) })
	eventstore.RegisterEventType(LabelPolicyChangedEventType, func() eventstore.Event { return new(LabelPolicyChangedEvent) })
	eventstore.RegisterEventType(LabelPolicyActivatedEventType, func() eventstore.Event { return new(LabelPolicyActivatedEvent) })
	eventstore.RegisterEventType(LabelPolicyLogoAddedEventType, func() eventstore.Event { return new(LabelPolicyLogoAddedEvent) })
	eventstore.RegisterEventType(LabelPolicyLogoRemovedEventType, func() eventstore.Event { return new(LabelPolicyLogoRemovedEvent) })
	eventstore.RegisterEventType(LabelPolicyIconAddedEventType, func() eventstore.Event { return new(LabelPolicyIconAddedEvent) })
	eventstore.RegisterEventType(LabelPolicyIconRemovedEventType, func() eventstore.Event { return new(LabelPolicyIconRemovedEvent) })
	eventstore.RegisterEventType(LabelPolicyLogoDarkAddedEventType, func() eventstore.Event { return new(LabelPolicyLogoDarkAddedEvent) })
	eventstore.RegisterEventType(LabelPolicyLogoDarkRemovedEventType, func() eventstore.Event { return new(LabelPolicyLogoDarkRemovedEvent) })
	eventstore.RegisterEventType(LabelPolicyIconDarkAddedEventType, func() eventstore.Event { return new(LabelPolicyIconDarkAddedEvent) })
	eventstore.RegisterEventType(LabelPolicyIconDarkRemovedEventType, func() eventstore.Event { return new(LabelPolicyIconDarkRemovedEvent) })
	eventstore.RegisterEventType(LabelPolicyFontAddedEventType, func() eventstore.Event { return new(LabelPolicyFontAddedEvent) })
	eventstore.RegisterEventType(LabelPolicyFontRemovedEventType, func() eventstore.Event { return new(LabelPolicyFontRemovedEvent) })
	eventstore.RegisterEventType(LabelPolicyAssetsRemovedEventType, func() eventstore.Event { return new(LabelPolicyAssetsRemovedEvent) })
	eventstore.RegisterEventType(LoginPolicyAddedEventType, func() eventstore.Event { return new(LoginPolicyAddedEvent) })
	eventstore.RegisterEventType(LoginPolicyChangedEventType, func() eventstore.Event { return new(LoginPolicyChangedEvent) })
	eventstore.RegisterEventType(OrgIAMPolicyAddedEventType, func() eventstore.Event { return new(OrgIAMPolicyAddedEvent) })
	eventstore.RegisterEventType(OrgIAMPolicyChangedEventType, func() eventstore.Event { return new(OrgIAMPolicyChangedEvent) })
	eventstore.RegisterEventType(PasswordAgePolicyAddedEventType, func() eventstore.Event { return new(PasswordAgePolicyAddedEvent) })
	eventstore.RegisterEventType(PasswordAgePolicyChangedEventType, func() eventstore.Event { return new(PasswordAgePolicyChangedEvent) })
	eventstore.RegisterEventType(PasswordComplexityPolicyAddedEventType, func() eventstore.Event { return new(PasswordComplexityPolicyAddedEvent) })
	eventstore.RegisterEventType(PasswordComplexityPolicyChangedEventType, func() eventstore.Event { return new(PasswordComplexityPolicyChangedEvent) })
	eventstore.RegisterEventType(LockoutPolicyAddedEventType, func() eventstore.Event { return new(LockoutPolicyAddedEvent) })
	eventstore.RegisterEventType(LockoutPolicyChangedEventType, func() eventstore.Event { return new(LockoutPolicyChangedEvent) })
	eventstore.RegisterEventType(PrivacyPolicyAddedEventType, func() eventstore.Event { return new(PrivacyPolicyAddedEvent) })
	eventstore.RegisterEventType(PrivacyPolicyChangedEventType, func() eventstore.Event { return new(PrivacyPolicyChangedEvent) })
	eventstore.RegisterEventType(MemberAddedEventType, func() eventstore.Event { return new(MemberAddedEvent) })
	eventstore.RegisterEventType(MemberChangedEventType, func() eventstore.Event { return new(MemberChangedEvent) })
	eventstore.RegisterEventType(MemberRemovedEventType, func() eventstore.Event { return new(MemberRemovedEvent) })
	eventstore.RegisterEventType(MemberCascadeRemovedEventType, func() eventstore.Event { return new(MemberCascadeRemovedEvent) })
	eventstore.RegisterEventType(IDPConfigAddedEventType, func() eventstore.Event { return new(IDPConfigAddedEvent) })
	eventstore.RegisterEventType(IDPConfigChangedEventType, func() eventstore.Event { return new(IDPConfigChangedEvent) })
	eventstore.RegisterEventType(IDPConfigRemovedEventType, func() eventstore.Event { return new(IDPConfigRemovedEvent) })
	eventstore.RegisterEventType(IDPConfigDeactivatedEventType, func() eventstore.Event { return new(IDPConfigDeactivatedEvent) })
	eventstore.RegisterEventType(IDPConfigReactivatedEventType, func() eventstore.Event { return new(IDPConfigReactivatedEvent) })
	eventstore.RegisterEventType(IDPOIDCConfigAddedEventType, func() eventstore.Event { return new(IDPOIDCConfigAddedEvent) })
	eventstore.RegisterEventType(IDPOIDCConfigChangedEventType, func() eventstore.Event { return new(IDPOIDCConfigChangedEvent) })
	eventstore.RegisterEventType(IDPJWTConfigAddedEventType, func() eventstore.Event { return new(IDPJWTConfigAddedEvent) })
	eventstore.RegisterEventType(IDPJWTConfigChangedEventType, func() eventstore.Event { return new(IDPJWTConfigChangedEvent) })
	eventstore.RegisterEventType(LoginPolicyIDPProviderAddedEventType, func() eventstore.Event { return new(IdentityProviderAddedEvent) })
	eventstore.RegisterEventType(LoginPolicyIDPProviderRemovedEventType, func() eventstore.Event { return new(IdentityProviderRemovedEvent) })
	eventstore.RegisterEventType(LoginPolicyIDPProviderCascadeRemovedEventType, func() eventstore.Event { return new(IdentityProviderCascadeRemovedEvent) })
	eventstore.RegisterEventType(LoginPolicySecondFactorAddedEventType, func() eventstore.Event { return new(LoginPolicySecondFactorAddedEvent) })
	eventstore.RegisterEventType(LoginPolicySecondFactorRemovedEventType, func() eventstore.Event { return new(LoginPolicySecondFactorRemovedEvent) })
	eventstore.RegisterEventType(LoginPolicyMultiFactorAddedEventType, func() eventstore.Event { return new(LoginPolicyMultiFactorAddedEvent) })
	eventstore.RegisterEventType(LoginPolicyMultiFactorRemovedEventType, func() eventstore.Event { return new(LoginPolicyMultiFactorRemovedEvent) })
	eventstore.RegisterEventType(MailTemplateAddedEventType, func() eventstore.Event { return new(MailTemplateAddedEvent) })
	eventstore.RegisterEventType(MailTemplateChangedEventType, func() eventstore.Event { return new(MailTemplateChangedEvent) })
	eventstore.RegisterEventType(MailTextAddedEventType, func() eventstore.Event { return new(MailTextAddedEvent) })
	eventstore.RegisterEventType(MailTextChangedEventType, func() eventstore.Event { return new(MailTextChangedEvent) })
	eventstore.RegisterEventType(CustomTextSetEventType, func() eventstore.Event { return new(CustomTextSetEvent) })
	eventstore.RegisterEventType(CustomTextRemovedEventType, func() eventstore.Event { return new(CustomTextRemovedEvent) })
	eventstore.RegisterEventType(CustomTextTemplateRemovedEventType, func() eventstore.Event { return new(CustomTextTemplateRemovedEvent) })
	eventstore.RegisterEventType(FeaturesSetEventType, func() eventstore.Event { return new(FeaturesSetEvent) })
}
