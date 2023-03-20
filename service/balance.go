package pave

import (
	"context"

	"encore.app/temporal"
	"encore.app/temporal/workflow"
	"encore.dev/rlog"
	"go.temporal.io/sdk/client"
)

// BalanceResponse returns id, available, reserved amount of the accont
type BalanceResponse struct {
    ID int
    Available int
    Reserved int
}

 // BalanceParams are the parameters for getting the balance
 type BalanceParams struct {
    ID int `json:"id"`
}

//encore:api public path=/balance
func (s *Service) Balance(ctx context.Context, p *BalanceParams) (*BalanceResponse, error) {
    options := client.StartWorkflowOptions{
        ID:        "pave-workflow",
        TaskQueue: paveTaskQueue,
    }
    we, err := s.client.ExecuteWorkflow(ctx, options, workflow.Balance, p.ID)
    if err != nil {
        return nil, err
    }
    rlog.Info("started workflow", "id", we.GetID(), "run_id", we.GetRunID())

    var result temporal.Account
    err = we.Get(ctx, &result)
    if err != nil {
        return nil, err
    }
    return &BalanceResponse{
        ID: result.ID,
        Available: result.Available,
        Reserved: result.Reserved,
    }, nil
}
