package apps

import "github.com/RandySteven/go-kopi/interfaces/usecases"

type Usecases struct {
	UserUsecase usecases.IUserUsecase
}

func NewUsecases() *Usecases {
	return &Usecases{}
}
