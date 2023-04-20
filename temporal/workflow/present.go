package workflow

import (
	"go.temporal.io/sdk/workflow"
)

func Present(ctx workflow.Context, workflowID string) (string, error) {

    signal := true
    if workflowID == "" {
        return "", nil
    }
    err :=  workflow.SignalExternalWorkflow(ctx, workflowID, "", AcceptSubmissionSignalName, signal).Get(ctx, nil)
    if err != nil {
        return "", nil
    }
    return "", nil
}

