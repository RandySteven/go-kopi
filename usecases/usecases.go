package usecases

import usecases_interfaces "github.com/RandySteven/go-kopi/interfaces/usecases"

type Usecases struct {
	UserUsecase usecases_interfaces.UserUsecase
}

func NewUsecases() *Usecases {
	return &Usecases{}
}
