package consumers

import (
	"context"
	"log"

	"github.com/nsqio/go-nsq"
)

type (
	DummyConsumer interface {
		Consumer
	}

	dummyConsumer struct {
	}
)

// DummyFunc implements [DummyConsumer].
func (d *dummyConsumer) Consume(ctx context.Context, nsqMessage *nsq.Message) error {
	log.Println(nsqMessage.Body)
	return nil
}

var _ DummyConsumer = &dummyConsumer{}

func NewDummyConsumer() DummyConsumer {
	return &dummyConsumer{}
}
