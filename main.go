package main

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/temporalio/temporalite"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	log2 "go.temporal.io/server/common/log"
)

var log = zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().Timestamp().Logger()

func main() {

	// Start temporalite server
	logger := log2.NewNoopLogger()
	server, err := temporalite.NewServer(temporalite.WithLogger(logger))
	if err != nil {
		panic(err)
	}

	if err := server.Start(); err != nil {
		panic(err)
	}
	c, err := server.NewClient(context.Background(), "default")
	if err != nil {
		panic(err)
	}
	// Start worker
	w := generateWorker(c)
	if err := w.Start(); err != nil {
		panic(err)
	}
	err = scheduleWorkflow(c)

	// At this point we have a server running in the background,
	// a worker running in the background and a WORKFLOW running in
	// the background.

	// Wait for the worker interrupt channel. A signal will be sent
	// from the local activity
	select {
	case <-worker.InterruptCh():
		// Handle the signal and gracefully stop the worker
		w.Stop()

	}

	// Worker has been stopped, print the status of the workflow
	// After stopping, retrieve the workflow
	wfRun := c.GetWorkflow(context.Background(), "Test", "")
	// Describe it
	describe, err := c.DescribeWorkflowExecution(context.Background(), wfRun.GetID(), wfRun.GetRunID())
	if err != nil {
		panic(err)
	}
	log.Warn().Msgf("The status for the workflow (wfId:%s,runId:%s) is: %s", describe.GetWorkflowExecutionInfo().GetExecution().GetWorkflowId(), describe.GetWorkflowExecutionInfo().GetExecution().GetRunId(), describe.GetWorkflowExecutionInfo().GetStatus())
}

func scheduleWorkflow(c client.Client) error {
	opts := client.StartWorkflowOptions{
		ID:                       "Test",
		TaskQueue:                "worker",
		WorkflowExecutionTimeout: 10 * time.Second,
	}
	log.Warn().Msg("Starting the workflow")
	_, err := c.ExecuteWorkflow(context.Background(), opts, "MyWorkflow")
	return err
}

func generateWorker(c client.Client) worker.Worker {

	tWorker := worker.New(c, "worker", worker.Options{
		WorkerStopTimeout: 10 * time.Second,
	},
	)

	tWorker.RegisterWorkflowWithOptions(MyWorkflow, workflow.RegisterOptions{
		Name: "MyWorkflow",
	})
	tWorker.RegisterActivityWithOptions(MyActivity, activity.RegisterOptions{
		Name: "MyActivity",
	})

	return tWorker

}
