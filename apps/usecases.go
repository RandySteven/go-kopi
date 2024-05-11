package apps

import "go_framework_dev/interfaces/usecases"

type Usecases struct {
	UserUsecase usecases.IUserUsecase
}

func NewUsecases() *Usecases {
	return &Usecases{}
}
