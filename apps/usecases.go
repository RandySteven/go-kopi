package apps

import (
	"github.com/RandySteven/go-kopi/interfaces/usecases"
	"github.com/RandySteven/go-kopi/pkg/db"
)

type Usecases struct {
	UserUsecase usecases.IUserUsecase
}

func newUsecases(repo *db.Repositories, services *Services) *Usecases {

	return &Usecases{}
}
