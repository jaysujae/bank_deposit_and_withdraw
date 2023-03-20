package workflow

import (
	"time"

	"encore.app/temporal"
	"encore.app/temporal/activity"
	"github.com/google/uuid"

	"go.temporal.io/api/enums/v1"
	"go.temporal.io/sdk/workflow"
)
var a *activity.Activities

const (
	AcceptSubmissionSignalName = "accept-submission"
	AcceptGracePeriod          = time.Second * 10
)


func Authorize(ctx workflow.Context, param temporal.AuthorizeParams) (temporal.AuthorizeResponse, error) {
    options := workflow.ActivityOptions{
        StartToCloseTimeout: time.Second * 5,
    }

    ctx = workflow.WithActivityOptions(ctx, options)

    var result string
    err := workflow.ExecuteActivity(ctx, a.CreateAuthorizeTransfer, param).Get(ctx, &result)
    if err != nil {
        return temporal.AuthorizeResponse{}, err
    }
	childId := uuid.NewString()[:5]
	childWorkflowOptions := workflow.ChildWorkflowOptions{
		ParentClosePolicy: enums.PARENT_CLOSE_POLICY_ABANDON,
		WorkflowID: childId,
	}
	ctx = workflow.WithChildOptions(ctx, childWorkflowOptions)
	childWorkflow := workflow.ExecuteChildWorkflow(ctx, Wait, result)
	_ = childWorkflow.GetChildWorkflowExecution().Get(ctx, nil)

    return temporal.AuthorizeResponse{WorkflowID: childId}, err
}