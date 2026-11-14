package handlers

import (
	"github.com/RandySteven/jlpt-exercise-be/topics"
	"github.com/RandySteven/jlpt-exercise-be/usecases"
)

type Handlers struct {
	UserHandler       IUserHandler
	OnboardingHandler IOnboardingHandler
	DummyHandler      IDummyHandler
}

func NewHandlers(usecases *usecases.Usecases, topics *topics.Topics) *Handlers {
	return &Handlers{
		UserHandler:       NewUserHandler(usecases.UserUsecase),
		OnboardingHandler: NewOnboardingHandler(usecases.OnboardingUsecase),
		DummyHandler:      NewDummyHandler(topics.DummyTopic),
	}
}
