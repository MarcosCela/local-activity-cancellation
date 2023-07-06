package main

import (
	"go.temporal.io/sdk/workflow"
)

// YourSimpleWorkflowDefintiion is the most basic Workflow Defintion.
func MyWorkflow(ctx workflow.Context) error {
	// Sync local activity
	workflow.GetLogger(ctx).Warn("Starting the local activity, wish me luck")
	ctx = workflow.WithLocalActivityOptions(ctx, workflow.LocalActivityOptions{
		// Higher timeout that the activity
		ScheduleToCloseTimeout: activityScheduleToCloseTimeout,
	})
	return workflow.ExecuteLocalActivity(ctx, "MyActivity").Get(ctx, nil)
}
