package main

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// YourSimpleWorkflowDefintiion is the most basic Workflow Defintion.
func MyWorkflow(ctx workflow.Context) error {
	// Sync local activity
	workflow.GetLogger(ctx).Warn("Starting the local activity, wish me luck")
	ctx = workflow.WithLocalActivityOptions(ctx, workflow.LocalActivityOptions{
		ScheduleToCloseTimeout: 5 * time.Second,
	})
	return workflow.ExecuteLocalActivity(ctx, "MyActivity").Get(ctx, nil)
}
