package handlers

import (
	"github.com/RandySteven/go-kopi/interfaces/handlers"
	"github.com/RandySteven/go-kopi/interfaces/usecases"
	"net/http"
)

type UserHandler struct {
	userUsecase usecases.IUserUsecase
}

func (u *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
}

func (u *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
}

func NewUserHandler(userUsecase usecases.IUserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

var _ handlers.IUserHandler = &UserHandler{}
