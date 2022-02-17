package internal

import (
	models2 "auth/internal/models"
)

type Repo interface {
	CreateUser(user *models2.User) (*models2.User, models2.ErrorCode)
	GetUser(user *models2.User) (*models2.User, models2.ErrorCode)
}

type UseCase interface {
	SignUp(user *models2.User) (*models2.User, models2.ErrorCode)
	SignIn(user *models2.User) (*models2.User, models2.ErrorCode)
}

type SessionRepo interface {
	SessionOn(user models2.User) (string, error)
	SessionOff(user models2.User) error
}

type SessionUsecase interface {
	Login(user models2.User) (string, error)
	Logout(user models2.User) error
}
