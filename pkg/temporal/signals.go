package temporal_client

import (
	"github.com/RandySteven/go-kopi/configs"
	"go.temporal.io/sdk/workflow"
)

type (
	SignalHandler func(ctx workflow.Context, signalName string, result interface{})

	SignalConsumer interface {
		StartConsume(ctx workflow.Context, config *configs.Config) (err error)
		StopConsume(ctx workflow.Context) (err error)
	}
)

//signalHandleFunc to validate the consumer signla function

//signalStream to consume the signal from the temporal and call signalHandlefunc