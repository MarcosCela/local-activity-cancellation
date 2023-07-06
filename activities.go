package main

import (
	"context"
	"syscall"
	"time"

	"go.temporal.io/sdk/activity"
)

func MyActivity(c context.Context) error {
	// Simulate that the worker is interrupted WHILE the activity is running
	activity.GetLogger(c).Warn("oh no we got terminated")
	syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	activity.GetLogger(c).Warn("the activity is still running tho...")
	time.Sleep(5 * time.Second)

	select {
	case <-c.Done():
		return c.Err()
	default:
		return nil
	}
}
