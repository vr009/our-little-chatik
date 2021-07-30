package auth

import "our-little-chatik/internal/models"

type database interface {
	CreateUser() error
	GetUser(User models.User) error
}
