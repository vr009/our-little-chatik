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

type Encrypter interface {
	EncryptString(stringToEncrypt, salt string) string
}
