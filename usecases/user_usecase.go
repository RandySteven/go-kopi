package usecases

import (
	"context"
	"go_framework_dev/apperror"
	"go_framework_dev/entities/payloads/requests"
	"go_framework_dev/entities/payloads/responses"
	"go_framework_dev/interfaces/usecases"
)

type userUsecase struct{}

func (u userUsecase) RegisterUser(ctx context.Context, request *requests.UserRegisterRequest) (response *responses.UserRegisterResponse, appError apperror.CustomError) {
	return
}

func (u userUsecase) LoginUser(ctx context.Context, request *requests.UserLoginRequest) (response *responses.UserLoginResponse, appError apperror.CustomError) {
	return
}

var _ usecases.IUserUsecase = &userUsecase{}
