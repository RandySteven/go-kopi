package handlers

import (
	"github.com/RandySteven/go-kopi/usecases"
)

type Handlers struct {
	UserHandler IUserHandler
}

func NewHandlers(usecases *usecases.Usecases) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(usecases.UserUsecase),
	}
}
