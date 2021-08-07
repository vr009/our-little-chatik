package auth

import "our-little-chatik/internal/models"

type Repo interface {
	CreateUser(User models.User) error
	GetUser(User models.User) (string, string, error)
}

type UseCase interface {
	SignIn(User models.User) (string, error)
	SignUp(User models.User) error
}
