package apps

import "go_framework_dev/interfaces/handlers"

type Handlers struct {
	UserHandler handlers.IUserHandler
}

func NewHandler(userHandler handlers.IUserHandler) *Handlers {
	return &Handlers{}
}
