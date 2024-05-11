package apps

import "github.com/RandySteven/go-kopi/interfaces/handlers"

type Handlers struct {
	UserHandler handlers.IUserHandler
}

func NewHandler(userHandler handlers.IUserHandler) *Handlers {
	return &Handlers{}
}
