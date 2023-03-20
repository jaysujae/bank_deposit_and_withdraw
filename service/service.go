package pave

import (
	"context"
	"fmt"

	"encore.app/temporal/activity"
	"encore.app/temporal/workflow"
	"encore.dev"
	tb "github.com/tigerbeetledb/tigerbeetle-go"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

// Use an environment-specific task queue so we can use the same
// Temporal Cluster for all cloud environments.
var (
    envName = encore.Meta().Environment.Name
    paveTaskQueue = envName + "-pave"
)

type DB interface {
	GetWorkflowIDByAmount(amount int) string
	InsertWorkflowID(amount int, workflowID string)
}

type inmemoryDB struct {
	m map[int][]string
}

func (d *inmemoryDB) GetWorkflowIDByAmount(amount int) string {
	if len(d.m[amount][0]) > 0{
		return d.m[amount][0]
	}
	return ""
}

func (d *inmemoryDB) InsertWorkflowID(amount int, workflowID string){
	d.m[amount] = append(d.m[amount], workflowID)
}

//encore:service
type Service struct {
	client client.Client
	worker worker.Worker
	DB DB
}

func initService() (*Service, error) {
	c, err := client.Dial(client.Options{})
	if err != nil {
		return nil, fmt.Errorf("create temporal client: %v", err)
	}
	client, err := tb.NewClient(0, []string{"3000"}, 1)

	w := worker.New(c, paveTaskQueue, worker.Options{})
	w.RegisterWorkflow(workflow.Balance)
	w.RegisterWorkflow(workflow.Authorize)
	w.RegisterWorkflow(workflow.Present)
	w.RegisterWorkflow(workflow.Wait)
	w.RegisterActivity(&activity.Activities{
		Client: client,
	})
	
	
	err = w.Start()
	if err != nil {
		c.Close()
		return nil, fmt.Errorf("start temporal worker: %v", err)
	}
	return &Service{
		client: c, 
		worker: w,
		DB: &inmemoryDB{
			m: map[int][]string{},
		}}, nil
}

func (s *Service) Shutdown(force context.Context) {
	s.client.Close()
	s.worker.Stop()
}