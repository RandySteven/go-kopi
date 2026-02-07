package temporal_client

import (
	"context"
	"time"
)

type (
	// Workflow represents the state and metadata of a Temporal workflow execution.
	// It contains identifiers, status, timing information, and arbitrary metadata.
	Workflow struct {
		ID         string                 // Internal unique identifier
		WorkflowID string                 // Temporal workflow ID
		RunID      string                 // Current run ID
		PrevRunID  string                 // Previous run ID (for continued workflows)
		Status     string                 // Current workflow status
		Metadata   map[string]interface{} // Additional workflow metadata
		StartTime  time.Time              // Workflow start timestamp
		EndTime    time.Time              // Workflow end timestamp (if completed)
	}

	// WorkflowExecution defines the interface for retrieving workflow execution results.
	WorkflowExecution interface {
		// GetWorkflowResponse retrieves the result of a completed workflow by its ID.
		GetWorkflowResponse(ctx context.Context, workflowID string) (response interface{}, err error)
	}
)
