package command

import (
	"context"
	"reflect"

	"github.com/caos/zitadel/internal/domain"
	caos_errs "github.com/caos/zitadel/internal/errors"
	"github.com/caos/zitadel/internal/repository/org"
)

func (c *Commands) ClearFlow(ctx context.Context, flowType domain.FlowType, resourceOwner string) (*domain.ObjectDetails, error) {
	if !flowType.Valid() || resourceOwner == "" {
		return nil, caos_errs.ThrowInvalidArgument(nil, "COMMAND-Dfw2h", "Errors.Flow.FlowTypeMissing")
	}
	existingFlow, err := c.getOrgFlowWriteModelByType(ctx, flowType, resourceOwner)
	if err != nil {
		return nil, err
	}
	if len(existingFlow.Triggers) == 0 {
		return nil, caos_errs.ThrowPreconditionFailed(nil, "COMMAND-DgGh3", "Errors.Flow.Empty")
	}
	orgAgg := OrgAggregateFromWriteModel(&existingFlow.WriteModel)
	pushedEvents, err := c.eventstore.Push(ctx, org.NewFlowClearedEvent(ctx, orgAgg, flowType))
	if err != nil {
		return nil, err
	}
	err = AppendAndReduce(existingFlow, pushedEvents...)
	if err != nil {
		return nil, err
	}
	return writeModelToObjectDetails(&existingFlow.WriteModel), nil
}

func (c *Commands) SetTriggerActions(ctx context.Context, flowType domain.FlowType, triggerType domain.TriggerType, actionIDs []string, resourceOwner string) (*domain.ObjectDetails, error) {
	if !flowType.Valid() || !triggerType.Valid() || resourceOwner == "" {
		return nil, caos_errs.ThrowInvalidArgument(nil, "COMMAND-Dfhj5", "Errors.Flow.FlowTypeMissing")
	}
	if !flowType.HasTrigger(triggerType) {
		return nil, caos_errs.ThrowInvalidArgument(nil, "COMMAND-Dfgh6", "Errors.Flow.WrongTriggerType")
	}
	existingFlow, err := c.getOrgFlowWriteModelByType(ctx, flowType, resourceOwner)
	if err != nil {
		return nil, err
	}
	if reflect.DeepEqual(existingFlow.Triggers[triggerType], actionIDs) {
		return nil, caos_errs.ThrowPreconditionFailed(nil, "COMMAND-Nfh52", "Errors.Flow.NoChanges")
	}
	if len(actionIDs) > 0 {
		exists, err := c.actionsIDsExist(ctx, actionIDs, resourceOwner)
		if err != nil {
			return nil, err
		}
		if !exists {
			return nil, caos_errs.ThrowPreconditionFailed(nil, "COMMAND-dg422", "Errors.Flow.ActionIDsNotExist")
		}
	}
	orgAgg := OrgAggregateFromWriteModel(&existingFlow.WriteModel)
	pushedEvents, err := c.eventstore.Push(ctx, org.NewTriggerActionsSetEvent(ctx, orgAgg, flowType, triggerType, actionIDs))
	if err != nil {
		return nil, err
	}
	err = AppendAndReduce(existingFlow, pushedEvents...)
	if err != nil {
		return nil, err
	}
	return writeModelToObjectDetails(&existingFlow.WriteModel), nil
}

func (c *Commands) getOrgFlowWriteModelByType(ctx context.Context, flowType domain.FlowType, resourceOwner string) (*OrgFlowWriteModel, error) {
	flowWriteModel := NewOrgFlowWriteModel(flowType, resourceOwner)
	err := c.eventstore.FilterToQueryReducer(ctx, flowWriteModel)
	if err != nil {
		return nil, err
	}
	return flowWriteModel, nil
}

func (c *Commands) actionsIDsExist(ctx context.Context, ids []string, resourceOwner string) (bool, error) {
	actionIDsModel := NewActionsExistModel(ids, resourceOwner)
	err := c.eventstore.FilterToQueryReducer(ctx, actionIDsModel)
	return len(actionIDsModel.actionIDs) == len(actionIDsModel.checkedIDs), err
}
