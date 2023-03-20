package pave

import (
	"context"

	"encore.app/temporal/workflow"
	"encore.dev/rlog"
	"go.temporal.io/sdk/client"
)

// PresentResponse returns avaialble and reserved ammount of the account
type PresentResponse struct {
    ID int
    Available int
    Reserved int
}

// PresentParams gets amount of present money, and find the same amount of authorize transaction with the oldest data and present it.
type PresentParams struct {
    Amount int
}

//encore:api public path=/present
func (s *Service) Present(ctx context.Context, p *PresentParams) (*PresentResponse, error) {
    options := client.StartWorkflowOptions{
        ID:        "pave-workflow",
        TaskQueue: paveTaskQueue,
    }
    we, err := s.client.ExecuteWorkflow(
        ctx, 
        options, 
        workflow.Present, 
        s.DB.GetWorkflowIDByAmount(p.Amount),
    )
    if err != nil {
        return nil, err
    }
    rlog.Info("started workflow", "id", we.GetID(), "run_id", we.GetRunID())
    // Get the results
    var result string
    err = we.Get(ctx, &result)
    if err != nil {
        return nil, err
    }
    return &PresentResponse{
        
    }, nil
}


