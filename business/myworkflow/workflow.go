package myworkflow

import (
	"fmt"
	"time"

	"github.com/muchlist/workflow_tempo/business/myactivity"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func MyWorkflow(ctx workflow.Context, name string) (string, error) {
	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        100 * time.Second, // 100 * initial interval
		MaximumAttempts:        0,                 // unlimited
		NonRetryableErrorTypes: []string{},        // empty
	}

	ao := workflow.ActivityOptions{
		ScheduleToCloseTimeout: 100 * time.Second,
		StartToCloseTimeout:    10 * time.Second,
		RetryPolicy:            retryPolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	err := workflow.ExecuteActivity(ctx, myactivity.Activity, name).Get(ctx, &result)
	if err != nil {
		return "", fmt.Errorf("execute activity: %w", err)
	}
	return result, nil
}
