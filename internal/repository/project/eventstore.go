package project

import (
	"github.com/caos/zitadel/internal/eventstore"
)

func RegisterEventMappers(es *eventstore.Eventstore) {
	es.RegisterFilterEventMapper(ProjectAddedType, ProjectAddedEventMapper).
		RegisterFilterEventMapper(ProjectChangedType, ProjectChangeEventMapper).
		RegisterFilterEventMapper(ProjectDeactivatedType, ProjectDeactivatedEventMapper).
		RegisterFilterEventMapper(ProjectReactivatedType, ProjectReactivatedEventMapper).
		RegisterFilterEventMapper(ProjectRemovedType, ProjectRemovedEventMapper).
		RegisterFilterEventMapper(MemberAddedType, MemberAddedEventMapper).
		RegisterFilterEventMapper(MemberChangedType, MemberChangedEventMapper).
		RegisterFilterEventMapper(MemberRemovedType, MemberRemovedEventMapper).
		RegisterFilterEventMapper(MemberCascadeRemovedType, MemberCascadeRemovedEventMapper).
		RegisterFilterEventMapper(RoleAddedType, RoleAddedEventMapper).
		RegisterFilterEventMapper(RoleChangedType, RoleChangedEventMapper).
		RegisterFilterEventMapper(RoleRemovedType, RoleRemovedEventMapper).
		RegisterFilterEventMapper(GrantAddedType, GrantAddedEventMapper).
		RegisterFilterEventMapper(GrantChangedType, GrantChangedEventMapper).
		RegisterFilterEventMapper(GrantCascadeChangedType, GrantCascadeChangedEventMapper).
		RegisterFilterEventMapper(GrantDeactivatedType, GrantDeactivateEventMapper).
		RegisterFilterEventMapper(GrantReactivatedType, GrantReactivatedEventMapper).
		RegisterFilterEventMapper(GrantRemovedType, GrantRemovedEventMapper).
		RegisterFilterEventMapper(GrantMemberAddedType, GrantMemberAddedEventMapper).
		RegisterFilterEventMapper(GrantMemberChangedType, GrantMemberChangedEventMapper).
		RegisterFilterEventMapper(GrantMemberRemovedType, GrantMemberRemovedEventMapper).
		RegisterFilterEventMapper(GrantMemberCascadeRemovedType, GrantMemberCascadeRemovedEventMapper).
		RegisterFilterEventMapper(ApplicationAddedType, ApplicationAddedEventMapper).
		RegisterFilterEventMapper(ApplicationChangedType, ApplicationChangedEventMapper).
		RegisterFilterEventMapper(ApplicationRemovedType, ApplicationRemovedEventMapper).
		RegisterFilterEventMapper(ApplicationDeactivatedType, ApplicationDeactivatedEventMapper).
		RegisterFilterEventMapper(ApplicationReactivatedType, ApplicationReactivatedEventMapper).
		RegisterFilterEventMapper(OIDCConfigAddedType, OIDCConfigAddedEventMapper).
		RegisterFilterEventMapper(OIDCConfigChangedType, OIDCConfigChangedEventMapper).
		RegisterFilterEventMapper(OIDCConfigSecretChangedType, OIDCConfigSecretChangedEventMapper).
		RegisterFilterEventMapper(OIDCClientSecretCheckSucceededType, OIDCConfigSecretCheckSucceededEventMapper).
		RegisterFilterEventMapper(OIDCClientSecretCheckFailedType, OIDCConfigSecretCheckFailedEventMapper).
		RegisterFilterEventMapper(APIConfigAddedType, APIConfigAddedEventMapper).
		RegisterFilterEventMapper(APIConfigChangedType, APIConfigChangedEventMapper).
		RegisterFilterEventMapper(APIConfigSecretChangedType, APIConfigSecretChangedEventMapper).
		RegisterFilterEventMapper(ApplicationKeyAddedEventType, ApplicationKeyAddedEventMapper).
		RegisterFilterEventMapper(ApplicationKeyRemovedEventType, ApplicationKeyRemovedEventMapper)

	eventstore.RegisterEventType(ProjectAddedType, func() eventstore.Event { return new(ProjectAddedEvent) })
	eventstore.RegisterEventType(ProjectChangedType, func() eventstore.Event { return new(ProjectChangeEvent) })
	eventstore.RegisterEventType(ProjectDeactivatedType, func() eventstore.Event { return new(ProjectDeactivatedEvent) })
	eventstore.RegisterEventType(ProjectReactivatedType, func() eventstore.Event { return new(ProjectReactivatedEvent) })
	eventstore.RegisterEventType(ProjectRemovedType, func() eventstore.Event { return new(ProjectRemovedEvent) })
	eventstore.RegisterEventType(MemberAddedType, func() eventstore.Event { return new(MemberAddedEvent) })
	eventstore.RegisterEventType(MemberChangedType, func() eventstore.Event { return new(MemberChangedEvent) })
	eventstore.RegisterEventType(MemberRemovedType, func() eventstore.Event { return new(MemberRemovedEvent) })
	eventstore.RegisterEventType(MemberCascadeRemovedType, func() eventstore.Event { return new(MemberCascadeRemovedEvent) })
	eventstore.RegisterEventType(RoleAddedType, func() eventstore.Event { return new(RoleAddedEvent) })
	eventstore.RegisterEventType(RoleChangedType, func() eventstore.Event { return new(RoleChangedEvent) })
	eventstore.RegisterEventType(RoleRemovedType, func() eventstore.Event { return new(RoleRemovedEvent) })
	eventstore.RegisterEventType(GrantAddedType, func() eventstore.Event { return new(GrantAddedEvent) })
	eventstore.RegisterEventType(GrantChangedType, func() eventstore.Event { return new(GrantChangedEvent) })
	eventstore.RegisterEventType(GrantCascadeChangedType, func() eventstore.Event { return new(GrantCascadeChangedEvent) })
	eventstore.RegisterEventType(GrantDeactivatedType, func() eventstore.Event { return new(GrantDeactivateEvent) })
	eventstore.RegisterEventType(GrantReactivatedType, func() eventstore.Event { return new(GrantReactivatedEvent) })
	eventstore.RegisterEventType(GrantRemovedType, func() eventstore.Event { return new(GrantRemovedEvent) })
	eventstore.RegisterEventType(GrantMemberAddedType, func() eventstore.Event { return new(GrantMemberAddedEvent) })
	eventstore.RegisterEventType(GrantMemberChangedType, func() eventstore.Event { return new(GrantMemberChangedEvent) })
	eventstore.RegisterEventType(GrantMemberRemovedType, func() eventstore.Event { return new(GrantMemberRemovedEvent) })
	eventstore.RegisterEventType(GrantMemberCascadeRemovedType, func() eventstore.Event { return new(GrantMemberCascadeRemovedEvent) })
	eventstore.RegisterEventType(ApplicationAddedType, func() eventstore.Event { return new(ApplicationAddedEvent) })
	eventstore.RegisterEventType(ApplicationChangedType, func() eventstore.Event { return new(ApplicationChangedEvent) })
	eventstore.RegisterEventType(ApplicationRemovedType, func() eventstore.Event { return new(ApplicationRemovedEvent) })
	eventstore.RegisterEventType(ApplicationDeactivatedType, func() eventstore.Event { return new(ApplicationDeactivatedEvent) })
	eventstore.RegisterEventType(ApplicationReactivatedType, func() eventstore.Event { return new(ApplicationReactivatedEvent) })
	eventstore.RegisterEventType(OIDCConfigAddedType, func() eventstore.Event { return new(OIDCConfigAddedEvent) })
	eventstore.RegisterEventType(OIDCConfigChangedType, func() eventstore.Event { return new(OIDCConfigChangedEvent) })
	eventstore.RegisterEventType(OIDCConfigSecretChangedType, func() eventstore.Event { return new(OIDCConfigSecretChangedEvent) })
	eventstore.RegisterEventType(OIDCClientSecretCheckSucceededType, func() eventstore.Event { return new(OIDCConfigSecretCheckSucceededEvent) })
	eventstore.RegisterEventType(OIDCClientSecretCheckFailedType, func() eventstore.Event { return new(OIDCConfigSecretCheckFailedEvent) })
	eventstore.RegisterEventType(APIConfigAddedType, func() eventstore.Event { return new(APIConfigAddedEvent) })
	eventstore.RegisterEventType(APIConfigChangedType, func() eventstore.Event { return new(APIConfigChangedEvent) })
	eventstore.RegisterEventType(APIConfigSecretChangedType, func() eventstore.Event { return new(APIConfigSecretChangedEvent) })
	eventstore.RegisterEventType(ApplicationKeyAddedEventType, func() eventstore.Event { return new(ApplicationKeyAddedEvent) })
	eventstore.RegisterEventType(ApplicationKeyRemovedEventType, func() eventstore.Event { return new(ApplicationKeyRemovedEvent) })
}
