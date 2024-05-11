package usecases

import (
	"context"
	"go_framework_dev/apperror"
	"go_framework_dev/entities/payloads/requests"
	"go_framework_dev/entities/payloads/responses"
)

type IUserUsecase interface {
	RegisterUser(ctx context.Context, request *requests.UserRegisterRequest) (response *responses.UserRegisterResponse, appError apperror.CustomError)
	LoginUser(ctx context.Context, request *requests.UserLoginRequest) (response *responses.UserLoginResponse, appError apperror.CustomError)
}
