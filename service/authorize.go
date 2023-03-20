package pave

import (
	"context"

	"encore.app/temporal"
	"encore.app/temporal/workflow"
	"encore.dev/rlog"
	"go.temporal.io/sdk/client"
)
type AuthorizeResponse struct {
    ID int
    Available int
    Reserved int
}

 type AuthorizeParams struct {
    ID int `json:"id"`
    Amount int
}
//encore:api public path=/authorize
func (s *Service) Authorize(ctx context.Context, p *AuthorizeParams) (*AuthorizeResponse, error) {
    options := client.StartWorkflowOptions{
        TaskQueue: paveTaskQueue,
    }
    we, err := s.client.ExecuteWorkflow(ctx, options, workflow.Authorize, temporal.AuthorizeParams{
        ID: p.ID,
        Amount: p.Amount,
    })
    if err != nil {
        return nil, err
    }
    rlog.Info("started workflow", "id", we.GetID(), "run_id", we.GetRunID())

    // Get the results
    var result temporal.AuthorizeResponse
    err = we.Get(ctx, &result)
    if err != nil {
        return nil, err
    }
    s.DB.InsertWorkflowID(p.Amount,result.WorkflowID)
    
    return &AuthorizeResponse{
    }, nil
}


