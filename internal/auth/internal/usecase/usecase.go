package usecase

import (
	"auth/internal"
	models2 "auth/internal/models"
)

type AuthUseCase struct {
	repo internal.Repo
}

func NewAuthUseCase(userRepo internal.Repo) *AuthUseCase {
	return &AuthUseCase{
		repo: userRepo,
	}
}

func (au *AuthUseCase) SignIn(User *models2.User) (*models2.User, models2.ErrorCode) {
	return au.repo.GetUser(User)
}
func (au *AuthUseCase) SignUp(user *models2.User) (*models2.User, models2.ErrorCode) {
	return au.repo.CreateUser(user)
}
