package consumer_interfaces

import "context"

type UserConsumer interface {
	RegisterUser(ctx context.Context) error
}
