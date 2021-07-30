package usecase

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"our-little-chatik/internal/auth"
	"our-little-chatik/internal/models"
	"time"
)

type AuthUseCase struct {
	repo           auth.Repo
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo *auth.Repo,
	salt string,
	key []byte,
	dur time.Duration) *AuthUseCase {
	return &AuthUseCase{
		repo:           *userRepo,
		hashSalt:       salt,
		signingKey:     key,
		expireDuration: dur,
	}
}

func (a *AuthUseCase) SignUp(username, password string) error {
	pswd := sha256.New()
	pswd.Write([]byte(password))
	pswd.Write([]byte(a.hashSalt))

	user := models.User{
		UserName: username,
		Password: fmt.Sprintf("%x", pswd.Sum(nil)),
	}

	return a.repo.CreateUser(user)
}

func (a *AuthUseCase) SignIn(username, password string) error {
	pswd := sha256.New()
	pswd.Write([]byte(password))
	pswd.Write([]byte(a.hashSalt))

	compStr := fmt.Sprintf("%x", pswd.Sum(nil))

	user := models.User{
		UserName: username,
		Password: fmt.Sprintf("%x", pswd.Sum(nil)),
	}

	if DBpswd, err := a.repo.GetUser(user); err != nil {
		return err
	} else if DBpswd != compStr {
		return errors.New(fmt.Sprintf("not correct password\n"))
	}

	return nil
}
