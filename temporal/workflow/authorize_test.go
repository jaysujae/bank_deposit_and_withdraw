package workflow

import (
	"testing"

	"encore.app/temporal"
	"encore.app/temporal/activity"
	"github.com/stretchr/testify/mock"

	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)


func Test_Authorize(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()
	env.RegisterWorkflow(Wait)

	var a *activity.Activities
	env.OnActivity(a.CreateAuthorizeTransfer, mock.Anything, temporal.AuthorizeParams{
		ID: 1,
		Amount: 100,
	}).Return("abcde", nil)
	env.OnWorkflow(Wait,mock.Anything, "abcde").Return(temporal.Account{})

	env.ExecuteWorkflow(Authorize, temporal.AuthorizeParams{
		ID: 1,
		Amount: 100,
	})
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())
}