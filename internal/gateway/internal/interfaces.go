package internal

import "our-little-chatik/internal/models"

type ServerUsecase interface {
	SignUp(user *models.User) (*models.User, models.ErrorCode)
	SignIn(user *models.User) (*models.User, models.ErrorCode)

	GetSessionID() error
	AddSession() error
	RemoveSession() error
}
