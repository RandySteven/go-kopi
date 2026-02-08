package handlers_rest

import (
	rest_interfaces "github.com/RandySteven/go-kopi/interfaces/handlers/rest"
	"github.com/RandySteven/go-kopi/usecases"
)

type Rests struct {
	UserRest rest_interfaces.UserRest
}

func NewRESTs(usecases *usecases.Usecases) *Rests {
	return &Rests{
		UserRest: NewUserRest(usecases.UserUsecase),
	}
}
