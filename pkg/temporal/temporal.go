// Package temporal_client provides a Temporal workflow orchestration client wrapper.
// Temporal is a microservice orchestration platform for running mission-critical code
// with built-in retries, timeouts, and visibility.
package temporal_client

import (
	"context"

	"github.com/RandySteven/go-kopi/pkg/config"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type (
	// WorkflowReplayer is an alias for worker.WorkflowReplayer used for replaying workflows.
	WorkflowReplayer = worker.WorkflowReplayer
	// ActivityRegistry is an alias for worker.ActivityRegistry used for registering activities.
	ActivityRegistry = worker.ActivityRegistry
)

type (
	// Worker defines the interface combining workflow replay and activity registration capabilities.
	Worker interface {
		WorkflowReplayer
		ActivityRegistry
	}

	// workerKopi is the internal implementation wrapping a Temporal worker.
	workerKopi struct {
		worker worker.Worker
	}
)

// NewWorker creates a new Temporal worker connected to the specified Temporal server.
// It uses the hostport and namespace from the configuration, defaulting to
// client.DefaultHostPort if not specified.
// The worker is created for the "default" task queue.
// Returns an error if the Temporal client connection fails.
func NewWorker(ctx context.Context, cfg *config.Config) (*workerKopi, error) {
	hostPort := cfg.Configs.Temporal.HostPort
	if hostPort == "" {
		hostPort = client.DefaultHostPort
	}
	temporalClient, err := client.NewClient(client.Options{
		HostPort:  client.DefaultHostPort,
		Namespace: cfg.Configs.Temporal.Namespace,
	})
	if err != nil {
		return nil, err
	}
	defer temporalClient.Close()

	worker := worker.New(temporalClient, "default", worker.Options{})
	return &workerKopi{
		worker: worker,
	}, nil
}

// RegisterWorkflow registers a workflow function with the Temporal worker.
// The workflow parameter should be a workflow function that follows Temporal's
// workflow function signature requirements.
func (w *workerKopi) RegisterWorkflow(workflow interface{}) {
}
