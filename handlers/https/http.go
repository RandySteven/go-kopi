package api_http

import (
	"github.com/RandySteven/go-kopi/interfaces/handlers"
	"github.com/RandySteven/go-kopi/usecases"
)

type HTTPs struct {
	UserHTTP handlers.IUserHandler
}

func NewHTTPs(usecases *usecases.Usecases) *HTTPs {
	return &HTTPs{
		UserHTTP: NewUserHTTP(usecases.UserUsecase),
	}
}
