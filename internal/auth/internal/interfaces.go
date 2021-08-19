package internal

import (
	models2 "auth/internal/models"
)

type Repo interface {
	CreateUser(User models2.User) (string, error)
	GetUser(User models2.User) (string, string, error)
	GetAllUser() ([]models2.User, error)
}

type UseCase interface {
	SignIn(User models2.User) (string, error)
	SignUp(User models2.User) (string, error)
	FetchUsers() ([]models2.User, error)
}
