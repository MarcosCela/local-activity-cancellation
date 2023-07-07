package main

import (
	"context"
	"syscall"
	"time"

	"go.temporal.io/sdk/activity"
)

func MyActivity(c context.Context) error {
	// Simulate that the worker is interrupted WHILE the activity is running.
	// Under normal conditions, the SIGINT/SIGTERM is handled and the Stop() is called
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)

	// Simulate a long-running activity here that also
	// respects context cancellation
	select {
	case <-c.Done():
		activity.GetLogger(c).Error("The activity was cancelled")
		return c.Err()
	case <-time.After(activityTimer):
		activity.GetLogger(c).Error("The activity completed correctly")
		return nil
	}
}
