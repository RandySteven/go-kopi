package apps

import "context"

type Services struct{}

func NewServices(ctx context.Context) (*Services, error) {
	return &Services{}, nil
}
