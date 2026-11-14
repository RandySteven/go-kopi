package consumers

import (
	"context"

	"github.com/RandySteven/jlpt-exercise-be/caches"
	"github.com/RandySteven/jlpt-exercise-be/repositories"
	"github.com/RandySteven/jlpt-exercise-be/topics"
	"github.com/nsqio/go-nsq"
)

type (
	Consumers struct {
		DummyConsumer DummyConsumer
	}

	Consumer interface {
		Consume(ctx context.Context, nsqMessage *nsq.Message) error
	}

	RunConsumer map[string]ConsumerFunc
)

func (r *Runners) RegisterConsumer(topic string, fun ConsumerFunc) *Runners {
	r.RunConsumers[topic] = fun
	return r
}

func NewConsumers(
	repo *repositories.Repositories,
	cache *caches.Caches,
	topics *topics.Topics,
) *Consumers {
	return &Consumers{
		DummyConsumer: NewDummyConsumer(),
	}
}
