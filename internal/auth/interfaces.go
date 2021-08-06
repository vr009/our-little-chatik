package auth

import "our-little-chatik/internal/models"

type Repo interface {
	CreateUser(User models.User) error
	GetUser(User models.User) (string, string, error)
}

type UseCase interface {
	SignIn(username, password string) (string, error)
	SignUp(username, password string) error
}
