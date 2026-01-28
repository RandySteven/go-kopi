package api_http

import (
	"net/http"

	"github.com/RandySteven/go-kopi/interfaces/handlers"
	usecases_interfaces "github.com/RandySteven/go-kopi/interfaces/usecases"
)

type UserHTTP struct {
	userUsecase usecases_interfaces.UserUsecase
}

func (u *UserHTTP) RegisterUser(w http.ResponseWriter, r *http.Request) {
}

func (u *UserHTTP) LoginUser(w http.ResponseWriter, r *http.Request) {
}

func NewUserHTTP(userUsecase usecases_interfaces.UserUsecase) *UserHTTP {
	return &UserHTTP{
		userUsecase: userUsecase,
	}
}

var _ handlers.IUserHandler = &UserHTTP{}
