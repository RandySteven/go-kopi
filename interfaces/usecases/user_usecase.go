package usecases

import (
	"context"
	"github.com/RandySteven/go-kopi/apperror"
	"github.com/RandySteven/go-kopi/entities/payloads/requests"
	"github.com/RandySteven/go-kopi/entities/payloads/responses"
)

type IUserUsecase interface {
	RegisterUser(ctx context.Context, request *requests.UserRegisterRequest) (response *responses.UserRegisterResponse, appError apperror.CustomError)
	LoginUser(ctx context.Context, request *requests.UserLoginRequest) (response *responses.UserLoginResponse, appError apperror.CustomError)
}
