package apps

import (
	handlers2 "github.com/RandySteven/go-kopi/handlers"
	"github.com/RandySteven/go-kopi/interfaces/handlers"
	"github.com/RandySteven/go-kopi/pkg/db"
)

type Handlers struct {
	UserHandler handlers.IUserHandler
}

func NewHandlers(repo *db.Repositories, services *Services) *Handlers {
	usecases := newUsecases(repo, services)
	return &Handlers{
		UserHandler: handlers2.NewUserHandler(usecases.UserUsecase),
	}
}
