package workflow

import (
	"testing"

	"encore.app/temporal"
	"encore.app/temporal/activity"
	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)


func Test_BalanceWorkflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	var a *activity.Activities
	env.OnActivity(a.GetAccount, mock.Anything, 1).Return(temporal.Account{
		ID: 1,
		Available: 100,
		Reserved: 100,
	}, nil)

	env.ExecuteWorkflow(Balance, 1)
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
}