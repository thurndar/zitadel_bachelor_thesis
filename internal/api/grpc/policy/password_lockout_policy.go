package policy

import (
	"github.com/caos/zitadel/internal/api/grpc/object"
	"github.com/caos/zitadel/internal/query"
	policy_pb "github.com/caos/zitadel/pkg/grpc/policy"
)

func ModelLockoutPolicyToPb(policy *query.LockoutPolicy) *policy_pb.LockoutPolicy {
	return &policy_pb.LockoutPolicy{
		IsDefault:           policy.IsDefault,
		MaxPasswordAttempts: policy.MaxPasswordAttempts,
		Details: object.ToViewDetailsPb(
			policy.Sequence,
			policy.CreationDate,
			policy.ChangeDate,
			policy.ResourceOwner,
		),
	}
}
