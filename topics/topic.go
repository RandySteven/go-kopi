package topics

import (
	"context"

	nsq_client "github.com/RandySteven/go-cook/nsq"
)

type (
	Topic interface {
		WriteMessage(ctx context.Context, value string) (err error)
		ReadMessage(ctx context.Context) (value string, err error)
	}

	Topics struct {
	}
)

func NewTopics(nsq nsq_client.Nsq) *Topics {
	return &Topics{}
}
