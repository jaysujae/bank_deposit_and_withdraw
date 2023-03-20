package workflow

import (
	"time"

	"encore.app/temporal"

	"go.temporal.io/sdk/workflow"
)

func waitForSubmission(ctx workflow.Context) (bool, error) {
	var response bool
	var err error
    

	s := workflow.NewSelector(ctx)

	ch := workflow.GetSignalChannel(ctx, AcceptSubmissionSignalName)
	s.AddReceive(ch, func(c workflow.ReceiveChannel, more bool) {
		var submission bool
		c.Receive(ctx, &submission)
		response = true
	})
	s.AddFuture(workflow.NewTimer(ctx, AcceptGracePeriod), func(f workflow.Future) {
		err = f.Get(ctx, nil)
		response = false
	})

	s.Select(ctx)

	return response, err
}

func Wait(ctx workflow.Context, result string) (temporal.Account, error) {
	options := workflow.ActivityOptions{
        StartToCloseTimeout: time.Second * 200,
		ScheduleToCloseTimeout: time.Second * 200,
    }
	ctx = workflow.WithActivityOptions(ctx, options)
	
	success, err := waitForSubmission(ctx)
	if err != nil {
        return temporal.Account{}, err
    }
	if !success{
		err := workflow.ExecuteActivity(ctx, a.CancelAuthorize, result).Get(ctx, nil)
		if err != nil {
			return temporal.Account{}, err
		}
	} else {
		err := workflow.ExecuteActivity(ctx, a.Present, result).Get(ctx, nil)
		if err != nil {
			return temporal.Account{}, err
		}
	}
	
    return temporal.Account{}, err
}