package user

import (
	"github.com/caos/zitadel/internal/eventstore"
)

func RegisterEventMappers(es *eventstore.Eventstore) {
	es.RegisterFilterEventMapper(UserV1AddedType, HumanAddedEventMapper).
		RegisterFilterEventMapper(UserV1RegisteredType, HumanRegisteredEventMapper).
		RegisterFilterEventMapper(UserV1InitialCodeAddedType, HumanInitialCodeAddedEventMapper).
		RegisterFilterEventMapper(UserV1InitialCodeSentType, HumanInitialCodeSentEventMapper).
		RegisterFilterEventMapper(UserV1InitializedCheckSucceededType, HumanInitializedCheckSucceededEventMapper).
		RegisterFilterEventMapper(UserV1InitializedCheckFailedType, HumanInitializedCheckFailedEventMapper).
		RegisterFilterEventMapper(UserV1SignedOutType, HumanSignedOutEventMapper).
		RegisterFilterEventMapper(UserV1PasswordChangedType, HumanPasswordChangedEventMapper).
		RegisterFilterEventMapper(UserV1PasswordCodeAddedType, HumanPasswordCodeAddedEventMapper).
		RegisterFilterEventMapper(UserV1PasswordCodeSentType, HumanPasswordCodeSentEventMapper).
		RegisterFilterEventMapper(UserV1PasswordCheckSucceededType, HumanPasswordCheckSucceededEventMapper).
		RegisterFilterEventMapper(UserV1PasswordCheckFailedType, HumanPasswordCheckFailedEventMapper).
		RegisterFilterEventMapper(UserV1EmailChangedType, HumanEmailChangedEventMapper).
		RegisterFilterEventMapper(UserV1EmailVerifiedType, HumanEmailVerifiedEventMapper).
		RegisterFilterEventMapper(UserV1EmailVerificationFailedType, HumanEmailVerificationFailedEventMapper).
		RegisterFilterEventMapper(UserV1EmailCodeAddedType, HumanEmailCodeAddedEventMapper).
		RegisterFilterEventMapper(UserV1EmailCodeSentType, HumanEmailCodeSentEventMapper).
		RegisterFilterEventMapper(UserV1PhoneChangedType, HumanPhoneChangedEventMapper).
		RegisterFilterEventMapper(UserV1PhoneRemovedType, HumanPhoneRemovedEventMapper).
		RegisterFilterEventMapper(UserV1PhoneVerifiedType, HumanPhoneVerifiedEventMapper).
		RegisterFilterEventMapper(UserV1PhoneVerificationFailedType, HumanPhoneVerificationFailedEventMapper).
		RegisterFilterEventMapper(UserV1PhoneCodeAddedType, HumanPhoneCodeAddedEventMapper).
		RegisterFilterEventMapper(UserV1PhoneCodeSentType, HumanPhoneCodeSentEventMapper).
		RegisterFilterEventMapper(UserV1ProfileChangedType, HumanProfileChangedEventMapper).
		RegisterFilterEventMapper(UserV1AddressChangedType, HumanAddressChangedEventMapper).
		RegisterFilterEventMapper(UserV1MFAInitSkippedType, HumanMFAInitSkippedEventMapper).
		RegisterFilterEventMapper(UserV1MFAOTPAddedType, HumanOTPAddedEventMapper).
		RegisterFilterEventMapper(UserV1MFAOTPVerifiedType, HumanOTPVerifiedEventMapper).
		RegisterFilterEventMapper(UserV1MFAOTPRemovedType, HumanOTPRemovedEventMapper).
		RegisterFilterEventMapper(UserV1MFAOTPCheckSucceededType, HumanOTPCheckSucceededEventMapper).
		RegisterFilterEventMapper(UserV1MFAOTPCheckFailedType, HumanOTPCheckFailedEventMapper).
		RegisterFilterEventMapper(UserLockedType, UserLockedEventMapper).
		RegisterFilterEventMapper(UserUnlockedType, UserUnlockedEventMapper).
		RegisterFilterEventMapper(UserDeactivatedType, UserDeactivatedEventMapper).
		RegisterFilterEventMapper(UserReactivatedType, UserReactivatedEventMapper).
		RegisterFilterEventMapper(UserRemovedType, UserRemovedEventMapper).
		RegisterFilterEventMapper(UserTokenAddedType, UserTokenAddedEventMapper).
		RegisterFilterEventMapper(UserTokenRemovedType, UserTokenRemovedEventMapper).
		RegisterFilterEventMapper(UserDomainClaimedType, DomainClaimedEventMapper).
		RegisterFilterEventMapper(UserDomainClaimedSentType, DomainClaimedSentEventMapper).
		RegisterFilterEventMapper(UserUserNameChangedType, UsernameChangedEventMapper).
		RegisterFilterEventMapper(MetadataSetType, MetadataSetEventMapper).
		RegisterFilterEventMapper(MetadataRemovedType, MetadataRemovedEventMapper).
		RegisterFilterEventMapper(MetadataRemovedAllType, MetadataRemovedAllEventMapper).
		RegisterFilterEventMapper(HumanAddedType, HumanAddedEventMapper).
		RegisterFilterEventMapper(HumanRegisteredType, HumanRegisteredEventMapper).
		RegisterFilterEventMapper(HumanInitialCodeAddedType, HumanInitialCodeAddedEventMapper).
		RegisterFilterEventMapper(HumanInitialCodeSentType, HumanInitialCodeSentEventMapper).
		RegisterFilterEventMapper(HumanInitializedCheckSucceededType, HumanInitializedCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanInitializedCheckFailedType, HumanInitializedCheckFailedEventMapper).
		RegisterFilterEventMapper(HumanSignedOutType, HumanSignedOutEventMapper).
		RegisterFilterEventMapper(HumanPasswordChangedType, HumanPasswordChangedEventMapper).
		RegisterFilterEventMapper(HumanPasswordCodeAddedType, HumanPasswordCodeAddedEventMapper).
		RegisterFilterEventMapper(HumanPasswordCodeSentType, HumanPasswordCodeSentEventMapper).
		RegisterFilterEventMapper(HumanPasswordCheckSucceededType, HumanPasswordCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanPasswordCheckFailedType, HumanPasswordCheckFailedEventMapper).
		RegisterFilterEventMapper(UserIDPLinkAddedType, UserIDPLinkAddedEventMapper).
		RegisterFilterEventMapper(UserIDPLinkRemovedType, UserIDPLinkRemovedEventMapper).
		RegisterFilterEventMapper(UserIDPLinkCascadeRemovedType, UserIDPLinkCascadeRemovedEventMapper).
		RegisterFilterEventMapper(UserIDPLoginCheckSucceededType, UserIDPCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanEmailChangedType, HumanEmailChangedEventMapper).
		RegisterFilterEventMapper(HumanEmailVerifiedType, HumanEmailVerifiedEventMapper).
		RegisterFilterEventMapper(HumanEmailVerificationFailedType, HumanEmailVerificationFailedEventMapper).
		RegisterFilterEventMapper(HumanEmailCodeAddedType, HumanEmailCodeAddedEventMapper).
		RegisterFilterEventMapper(HumanEmailCodeSentType, HumanEmailCodeSentEventMapper).
		RegisterFilterEventMapper(HumanPhoneChangedType, HumanPhoneChangedEventMapper).
		RegisterFilterEventMapper(HumanPhoneRemovedType, HumanPhoneRemovedEventMapper).
		RegisterFilterEventMapper(HumanPhoneVerifiedType, HumanPhoneVerifiedEventMapper).
		RegisterFilterEventMapper(HumanPhoneVerificationFailedType, HumanPhoneVerificationFailedEventMapper).
		RegisterFilterEventMapper(HumanPhoneCodeAddedType, HumanPhoneCodeAddedEventMapper).
		RegisterFilterEventMapper(HumanPhoneCodeSentType, HumanPhoneCodeSentEventMapper).
		RegisterFilterEventMapper(HumanProfileChangedType, HumanProfileChangedEventMapper).
		RegisterFilterEventMapper(HumanAvatarAddedType, HumanAvatarAddedEventMapper).
		RegisterFilterEventMapper(HumanAvatarRemovedType, HumanAvatarRemovedEventMapper).
		RegisterFilterEventMapper(HumanAddressChangedType, HumanAddressChangedEventMapper).
		RegisterFilterEventMapper(HumanMFAInitSkippedType, HumanMFAInitSkippedEventMapper).
		RegisterFilterEventMapper(HumanMFAOTPAddedType, HumanOTPAddedEventMapper).
		RegisterFilterEventMapper(HumanMFAOTPVerifiedType, HumanOTPVerifiedEventMapper).
		RegisterFilterEventMapper(HumanMFAOTPRemovedType, HumanOTPRemovedEventMapper).
		RegisterFilterEventMapper(HumanMFAOTPCheckSucceededType, HumanOTPCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanMFAOTPCheckFailedType, HumanOTPCheckFailedEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenAddedType, HumanU2FAddedEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenVerifiedType, HumanU2FVerifiedEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenSignCountChangedType, HumanU2FSignCountChangedEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenRemovedType, HumanU2FRemovedEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenBeginLoginType, HumanU2FBeginLoginEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenCheckSucceededType, HumanU2FCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanU2FTokenCheckFailedType, HumanU2FCheckFailedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenAddedType, HumanPasswordlessAddedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenVerifiedType, HumanPasswordlessVerifiedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenSignCountChangedType, HumanPasswordlessSignCountChangedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenRemovedType, HumanPasswordlessRemovedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenBeginLoginType, HumanPasswordlessBeginLoginEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenCheckSucceededType, HumanPasswordlessCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessTokenCheckFailedType, HumanPasswordlessCheckFailedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessInitCodeAddedType, HumanPasswordlessInitCodeAddedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessInitCodeRequestedType, HumanPasswordlessInitCodeRequestedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessInitCodeSentType, HumanPasswordlessInitCodeSentEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessInitCodeCheckFailedType, HumanPasswordlessInitCodeCodeCheckFailedEventMapper).
		RegisterFilterEventMapper(HumanPasswordlessInitCodeCheckSucceededType, HumanPasswordlessInitCodeCodeCheckSucceededEventMapper).
		RegisterFilterEventMapper(HumanRefreshTokenAddedType, HumanRefreshTokenAddedEventMapper).
		RegisterFilterEventMapper(HumanRefreshTokenRenewedType, HumanRefreshTokenRenewedEventEventMapper).
		RegisterFilterEventMapper(HumanRefreshTokenRemovedType, HumanRefreshTokenRemovedEventEventMapper).
		RegisterFilterEventMapper(MachineAddedEventType, MachineAddedEventMapper).
		RegisterFilterEventMapper(MachineChangedEventType, MachineChangedEventMapper).
		RegisterFilterEventMapper(MachineKeyAddedEventType, MachineKeyAddedEventMapper).
		RegisterFilterEventMapper(MachineKeyRemovedEventType, MachineKeyRemovedEventMapper).
		RegisterFilterEventMapper(PersonalAccessTokenAddedType, PersonalAccessTokenAddedEventMapper).
		RegisterFilterEventMapper(PersonalAccessTokenRemovedType, PersonalAccessTokenRemovedEventMapper)

	eventstore.RegisterEventType(UserV1AddedType, func() eventstore.Event { return new(HumanAddedEvent) })
	eventstore.RegisterEventType(UserV1RegisteredType, func() eventstore.Event { return new(HumanRegisteredEvent) })
	eventstore.RegisterEventType(UserV1InitialCodeAddedType, func() eventstore.Event { return new(HumanInitialCodeAddedEvent) })
	eventstore.RegisterEventType(UserV1InitialCodeSentType, func() eventstore.Event { return new(HumanInitialCodeSentEvent) })
	eventstore.RegisterEventType(UserV1InitializedCheckSucceededType, func() eventstore.Event { return new(HumanInitializedCheckSucceededEvent) })
	eventstore.RegisterEventType(UserV1InitializedCheckFailedType, func() eventstore.Event { return new(HumanInitializedCheckFailedEvent) })
	eventstore.RegisterEventType(UserV1SignedOutType, func() eventstore.Event { return new(HumanSignedOutEvent) })
	eventstore.RegisterEventType(UserV1PasswordChangedType, func() eventstore.Event { return new(HumanPasswordChangedEvent) })
	eventstore.RegisterEventType(UserV1PasswordCodeAddedType, func() eventstore.Event { return new(HumanPasswordCodeAddedEvent) })
	eventstore.RegisterEventType(UserV1PasswordCodeSentType, func() eventstore.Event { return new(HumanPasswordCodeSentEvent) })
	eventstore.RegisterEventType(UserV1PasswordCheckSucceededType, func() eventstore.Event { return new(HumanPasswordCheckSucceededEvent) })
	eventstore.RegisterEventType(UserV1PasswordCheckFailedType, func() eventstore.Event { return new(HumanPasswordCheckFailedEvent) })
	eventstore.RegisterEventType(UserV1EmailChangedType, func() eventstore.Event { return new(HumanEmailChangedEvent) })
	eventstore.RegisterEventType(UserV1EmailVerifiedType, func() eventstore.Event { return new(HumanEmailVerifiedEvent) })
	eventstore.RegisterEventType(UserV1EmailVerificationFailedType, func() eventstore.Event { return new(HumanEmailVerificationFailedEvent) })
	eventstore.RegisterEventType(UserV1EmailCodeAddedType, func() eventstore.Event { return new(HumanEmailCodeAddedEvent) })
	eventstore.RegisterEventType(UserV1EmailCodeSentType, func() eventstore.Event { return new(HumanEmailCodeSentEvent) })
	eventstore.RegisterEventType(UserV1PhoneChangedType, func() eventstore.Event { return new(HumanPhoneChangedEvent) })
	eventstore.RegisterEventType(UserV1PhoneRemovedType, func() eventstore.Event { return new(HumanPhoneRemovedEvent) })
	eventstore.RegisterEventType(UserV1PhoneVerifiedType, func() eventstore.Event { return new(HumanPhoneVerifiedEvent) })
	eventstore.RegisterEventType(UserV1PhoneVerificationFailedType, func() eventstore.Event { return new(HumanPhoneVerificationFailedEvent) })
	eventstore.RegisterEventType(UserV1PhoneCodeAddedType, func() eventstore.Event { return new(HumanPhoneCodeAddedEvent) })
	eventstore.RegisterEventType(UserV1PhoneCodeSentType, func() eventstore.Event { return new(HumanPhoneCodeSentEvent) })
	eventstore.RegisterEventType(UserV1ProfileChangedType, func() eventstore.Event { return new(HumanProfileChangedEvent) })
	eventstore.RegisterEventType(UserV1AddressChangedType, func() eventstore.Event { return new(HumanAddressChangedEvent) })
	eventstore.RegisterEventType(UserV1MFAInitSkippedType, func() eventstore.Event { return new(HumanMFAInitSkippedEvent) })
	eventstore.RegisterEventType(UserV1MFAOTPAddedType, func() eventstore.Event { return new(HumanOTPAddedEvent) })
	eventstore.RegisterEventType(UserV1MFAOTPVerifiedType, func() eventstore.Event { return new(HumanOTPVerifiedEvent) })
	eventstore.RegisterEventType(UserV1MFAOTPRemovedType, func() eventstore.Event { return new(HumanOTPRemovedEvent) })
	eventstore.RegisterEventType(UserV1MFAOTPCheckSucceededType, func() eventstore.Event { return new(HumanOTPCheckSucceededEvent) })
	eventstore.RegisterEventType(UserV1MFAOTPCheckFailedType, func() eventstore.Event { return new(HumanOTPCheckFailedEvent) })
	eventstore.RegisterEventType(UserLockedType, func() eventstore.Event { return new(UserLockedEvent) })
	eventstore.RegisterEventType(UserUnlockedType, func() eventstore.Event { return new(UserUnlockedEvent) })
	eventstore.RegisterEventType(UserDeactivatedType, func() eventstore.Event { return new(UserDeactivatedEvent) })
	eventstore.RegisterEventType(UserReactivatedType, func() eventstore.Event { return new(UserReactivatedEvent) })
	eventstore.RegisterEventType(UserRemovedType, func() eventstore.Event { return new(UserRemovedEvent) })
	eventstore.RegisterEventType(UserTokenAddedType, func() eventstore.Event { return new(UserTokenAddedEvent) })
	eventstore.RegisterEventType(UserTokenRemovedType, func() eventstore.Event { return new(UserTokenRemovedEvent) })
	eventstore.RegisterEventType(UserDomainClaimedType, func() eventstore.Event { return new(DomainClaimedEvent) })
	eventstore.RegisterEventType(UserDomainClaimedSentType, func() eventstore.Event { return new(DomainClaimedSentEvent) })
	eventstore.RegisterEventType(UserUserNameChangedType, func() eventstore.Event { return new(UsernameChangedEvent) })
	eventstore.RegisterEventType(MetadataSetType, func() eventstore.Event { return new(MetadataSetEvent) })
	eventstore.RegisterEventType(MetadataRemovedType, func() eventstore.Event { return new(MetadataRemovedEvent) })
	eventstore.RegisterEventType(MetadataRemovedAllType, func() eventstore.Event { return new(MetadataRemovedAllEvent) })
	eventstore.RegisterEventType(HumanAddedType, func() eventstore.Event { return new(HumanAddedEvent) })
	eventstore.RegisterEventType(HumanRegisteredType, func() eventstore.Event { return new(HumanRegisteredEvent) })
	eventstore.RegisterEventType(HumanInitialCodeAddedType, func() eventstore.Event { return new(HumanInitialCodeAddedEvent) })
	eventstore.RegisterEventType(HumanInitialCodeSentType, func() eventstore.Event { return new(HumanInitialCodeSentEvent) })
	eventstore.RegisterEventType(HumanInitializedCheckSucceededType, func() eventstore.Event { return new(HumanInitializedCheckSucceededEvent) })
	eventstore.RegisterEventType(HumanInitializedCheckFailedType, func() eventstore.Event { return new(HumanInitializedCheckFailedEvent) })
	eventstore.RegisterEventType(HumanSignedOutType, func() eventstore.Event { return new(HumanSignedOutEvent) })
	eventstore.RegisterEventType(HumanPasswordChangedType, func() eventstore.Event { return new(HumanPasswordChangedEvent) })
	eventstore.RegisterEventType(HumanPasswordCodeAddedType, func() eventstore.Event { return new(HumanPasswordCodeAddedEvent) })
	eventstore.RegisterEventType(HumanPasswordCodeSentType, func() eventstore.Event { return new(HumanPasswordCodeSentEvent) })
	eventstore.RegisterEventType(HumanPasswordCheckSucceededType, func() eventstore.Event { return new(HumanPasswordCheckSucceededEvent) })
	eventstore.RegisterEventType(HumanPasswordCheckFailedType, func() eventstore.Event { return new(HumanPasswordCheckFailedEvent) })
	eventstore.RegisterEventType(UserIDPLinkAddedType, func() eventstore.Event { return new(UserIDPLinkAddedEvent) })
	eventstore.RegisterEventType(UserIDPLinkRemovedType, func() eventstore.Event { return new(UserIDPLinkRemovedEvent) })
	eventstore.RegisterEventType(UserIDPLinkCascadeRemovedType, func() eventstore.Event { return new(UserIDPLinkCascadeRemovedEvent) })
	eventstore.RegisterEventType(UserIDPLoginCheckSucceededType, func() eventstore.Event { return new(UserIDPCheckSucceededEvent) })
	eventstore.RegisterEventType(HumanEmailChangedType, func() eventstore.Event { return new(HumanEmailChangedEvent) })
	eventstore.RegisterEventType(HumanEmailVerifiedType, func() eventstore.Event { return new(HumanEmailVerifiedEvent) })
	eventstore.RegisterEventType(HumanEmailVerificationFailedType, func() eventstore.Event { return new(HumanEmailVerificationFailedEvent) })
	eventstore.RegisterEventType(HumanEmailCodeAddedType, func() eventstore.Event { return new(HumanEmailCodeAddedEvent) })
	eventstore.RegisterEventType(HumanEmailCodeSentType, func() eventstore.Event { return new(HumanEmailCodeSentEvent) })
	eventstore.RegisterEventType(HumanPhoneChangedType, func() eventstore.Event { return new(HumanPhoneChangedEvent) })
	eventstore.RegisterEventType(HumanPhoneRemovedType, func() eventstore.Event { return new(HumanPhoneRemovedEvent) })
	eventstore.RegisterEventType(HumanPhoneVerifiedType, func() eventstore.Event { return new(HumanPhoneVerifiedEvent) })
	eventstore.RegisterEventType(HumanPhoneVerificationFailedType, func() eventstore.Event { return new(HumanPhoneVerificationFailedEvent) })
	eventstore.RegisterEventType(HumanPhoneCodeAddedType, func() eventstore.Event { return new(HumanPhoneCodeAddedEvent) })
	eventstore.RegisterEventType(HumanPhoneCodeSentType, func() eventstore.Event { return new(HumanPhoneCodeSentEvent) })
	eventstore.RegisterEventType(HumanProfileChangedType, func() eventstore.Event { return new(HumanProfileChangedEvent) })
	eventstore.RegisterEventType(HumanAvatarAddedType, func() eventstore.Event { return new(HumanAvatarAddedEvent) })
	eventstore.RegisterEventType(HumanAvatarRemovedType, func() eventstore.Event { return new(HumanAvatarRemovedEvent) })
	eventstore.RegisterEventType(HumanAddressChangedType, func() eventstore.Event { return new(HumanAddressChangedEvent) })
	eventstore.RegisterEventType(HumanMFAInitSkippedType, func() eventstore.Event { return new(HumanMFAInitSkippedEvent) })
	eventstore.RegisterEventType(HumanMFAOTPAddedType, func() eventstore.Event { return new(HumanOTPAddedEvent) })
	eventstore.RegisterEventType(HumanMFAOTPVerifiedType, func() eventstore.Event { return new(HumanOTPVerifiedEvent) })
	eventstore.RegisterEventType(HumanMFAOTPRemovedType, func() eventstore.Event { return new(HumanOTPRemovedEvent) })
	eventstore.RegisterEventType(HumanMFAOTPCheckSucceededType, func() eventstore.Event { return new(HumanOTPCheckSucceededEvent) })
	eventstore.RegisterEventType(HumanMFAOTPCheckFailedType, func() eventstore.Event { return new(HumanOTPCheckFailedEvent) })
	eventstore.RegisterEventType(HumanU2FTokenAddedType, func() eventstore.Event { return new(HumanU2FAddedEvent) })
	eventstore.RegisterEventType(HumanU2FTokenVerifiedType, func() eventstore.Event { return new(HumanU2FVerifiedEvent) })
	eventstore.RegisterEventType(HumanU2FTokenSignCountChangedType, func() eventstore.Event { return new(HumanU2FSignCountChangedEvent) })
	eventstore.RegisterEventType(HumanU2FTokenRemovedType, func() eventstore.Event { return new(HumanU2FRemovedEvent) })
	eventstore.RegisterEventType(HumanU2FTokenBeginLoginType, func() eventstore.Event { return new(HumanU2FBeginLoginEvent) })
	eventstore.RegisterEventType(HumanU2FTokenCheckSucceededType, func() eventstore.Event { return new(HumanU2FCheckSucceededEvent) })
	eventstore.RegisterEventType(HumanU2FTokenCheckFailedType, func() eventstore.Event { return new(HumanU2FCheckFailedEvent) })
	eventstore.RegisterEventType(HumanPasswordlessTokenAddedType, func() eventstore.Event { return new(HumanPasswordlessAddedEvent) })
	eventstore.RegisterEventType(HumanPasswordlessTokenVerifiedType, func() eventstore.Event { return new(HumanPasswordlessVerifiedEvent) })
	eventstore.RegisterEventType(HumanPasswordlessTokenSignCountChangedType, func() eventstore.Event { return new(HumanPasswordlessSignCountChangedEvent) })
	eventstore.RegisterEventType(HumanPasswordlessTokenRemovedType, func() eventstore.Event { return new(HumanPasswordlessRemovedEvent) })
	eventstore.RegisterEventType(HumanPasswordlessTokenBeginLoginType, func() eventstore.Event { return new(HumanPasswordlessBeginLoginEvent) })
	eventstore.RegisterEventType(HumanPasswordlessTokenCheckSucceededType, func() eventstore.Event { return new(HumanPasswordlessCheckSucceededEvent) })
	eventstore.RegisterEventType(HumanPasswordlessTokenCheckFailedType, func() eventstore.Event { return new(HumanPasswordlessCheckFailedEvent) })
	eventstore.RegisterEventType(HumanPasswordlessInitCodeAddedType, func() eventstore.Event { return new(HumanPasswordlessInitCodeAddedEvent) })
	eventstore.RegisterEventType(HumanPasswordlessInitCodeRequestedType, func() eventstore.Event { return new(HumanPasswordlessInitCodeRequestedEvent) })
	eventstore.RegisterEventType(HumanPasswordlessInitCodeSentType, func() eventstore.Event { return new(HumanPasswordlessInitCodeSentEvent) })
	eventstore.RegisterEventType(HumanPasswordlessInitCodeCheckFailedType, func() eventstore.Event { return new(HumanPasswordlessInitCodeCheckFailedEvent) })
	eventstore.RegisterEventType(HumanPasswordlessInitCodeCheckSucceededType, func() eventstore.Event { return new(HumanPasswordlessInitCodeCheckSucceededEvent) })
	eventstore.RegisterEventType(HumanRefreshTokenAddedType, func() eventstore.Event { return new(HumanRefreshTokenAddedEvent) })
	eventstore.RegisterEventType(HumanRefreshTokenRenewedType, func() eventstore.Event { return new(HumanRefreshTokenRenewedEvent) })
	eventstore.RegisterEventType(HumanRefreshTokenRemovedType, func() eventstore.Event { return new(HumanRefreshTokenRemovedEvent) })
	eventstore.RegisterEventType(MachineAddedEventType, func() eventstore.Event { return new(MachineAddedEvent) })
	eventstore.RegisterEventType(MachineChangedEventType, func() eventstore.Event { return new(MachineChangedEvent) })
	eventstore.RegisterEventType(MachineKeyAddedEventType, func() eventstore.Event { return new(MachineKeyAddedEvent) })
	eventstore.RegisterEventType(MachineKeyRemovedEventType, func() eventstore.Event { return new(MachineKeyRemovedEvent) })
	eventstore.RegisterEventType(PersonalAccessTokenAddedType, func() eventstore.Event { return new(PersonalAccessTokenAddedEvent) })
	eventstore.RegisterEventType(PersonalAccessTokenRemovedType, func() eventstore.Event { return new(PersonalAccessTokenRemovedEvent) })
}
