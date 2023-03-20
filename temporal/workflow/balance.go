package workflow

import (
	"time"

	"encore.app/temporal"
	"go.temporal.io/sdk/workflow"
)

func Balance(ctx workflow.Context, accountNumber int) (temporal.Account, error) {
    options := workflow.ActivityOptions{
        StartToCloseTimeout: time.Second * 5,
    }

    ctx = workflow.WithActivityOptions(ctx, options)

    var result temporal.Account
    err := workflow.ExecuteActivity(ctx, a.GetAccount, accountNumber).Get(ctx, &result)
    if err != nil {
        return temporal.Account{}, err
    }
    return result, err
}
