package usecase

import (
	"auth/internal"
	models2 "auth/internal/models"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"log"
	"net/http"
	"time"
)

type AuthUseCase struct {
	repo           internal.Repo
	hashSalt       string
	signingKey     []byte
	expireDuration time.Duration
}

func NewAuthUseCase(
	userRepo internal.Repo,
	salt string,
	key []byte,
	dur time.Duration) *AuthUseCase {
	return &AuthUseCase{
		repo:           userRepo,
		hashSalt:       salt,
		signingKey:     key,
		expireDuration: dur,
	}
}

type Claims struct {
	jwt.StandardClaims
	UUID string
}

func ParseToken(AccesToken string, SigningKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(AccesToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return SigningKey, nil
	})

	if err != nil {
		log.Println(err)
		return "", err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.UUID, nil
	}

	return "", nil
}

func (a *AuthUseCase) SignUp(User models2.User) (string, error) {
	if User.Username == "" || User.Password == "" {
		return "", errors.New("bad")
	}
	pswd := sha256.New()
	pswd.Write([]byte(User.Password))
	pswd.Write([]byte(a.hashSalt))

	DBuser := models2.User{
		Firstname: User.Firstname,
		Lastname:  User.Lastname,
		Username:  User.Username,
		Password:  fmt.Sprintf("%x", pswd.Sum(nil)),
	}

	uuid, err := a.repo.CreateUser(DBuser)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
		UUID: uuid,
	})

	return token.SignedString(a.signingKey)
}

func (a *AuthUseCase) SignIn(User models2.User) (string, error) {

	if User.Username == "" || User.Password == "" {
		return "", errors.New("bad")
	}

	pswd := sha256.New()
	pswd.Write([]byte(User.Password))
	pswd.Write([]byte(a.hashSalt))

	compStr := fmt.Sprintf("%x", pswd.Sum(nil))

	user := models2.User{
		Username: User.Username,
		Password: fmt.Sprintf("%x", pswd.Sum(nil)),
	}

	uuid, DBpswd, err := a.repo.GetUser(user)

	if err != nil {
		return "", err
	} else if DBpswd != compStr {
		return "", errors.New(fmt.Sprintf("not correct password\n"))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
		UUID: uuid,
	})

	return token.SignedString(a.signingKey)
}

func (a *AuthUseCase) AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		token, err := r.Cookie("ssid")
		if err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}

		if userId, err := ParseToken(token.Value, a.signingKey); err != nil {
			http.Error(w, "Forbidden", http.StatusForbidden)
		} else {
			log.Println(userId)
			next.ServeHTTP(w, r)
		}
	})
}

func (a *AuthUseCase) FetchUsers() ([]models2.User, error) {
	return a.repo.GetAllUser()
}
