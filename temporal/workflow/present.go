package workflow

import (
	"go.temporal.io/sdk/workflow"
)

func Present(ctx workflow.Context, workflowID string) (string, error) {

    signal := true
    err :=  workflow.SignalExternalWorkflow(ctx, workflowID, "", AcceptSubmissionSignalName, signal).Get(ctx, nil)
    
    return "", err
}

