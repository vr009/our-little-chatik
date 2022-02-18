package repo

import (
	"auth/internal"
	models2 "auth/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"strings"
)

type PGRepo struct {
	service *pgxpool.Pool
	encr    internal.Encrypter
}

func NewPGRepo(service *pgxpool.Pool, encr internal.Encrypter) *PGRepo {
	return &PGRepo{
		service: service,
		encr:    encr,
	}
}

func (repo *PGRepo) CreateUser(user *models2.User) (*models2.User, models2.ErrorCode) {
	uuidWithHyphen := uuid.New()
	uuidstr := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	log.Println("new")
	str := "insert into users values ($1 ,$2, $3, $4, $5, $6)"

	salt := uuid.New().String()
	user.Password = repo.encr.EncryptString(user.Password, salt)

	if _, err := repo.service.Exec(context.Background(),
		str, uuidstr, user.Username, user.Password, user.Firstname, user.Lastname, salt); err != nil {
		log.Println(err)
		return nil, models2.EXISTS
	}

	user.Uuid = uuidstr
	user.Password = ""
	return user, models2.OK
}

func (repo *PGRepo) GetUser(user *models2.User) (*models2.User, models2.ErrorCode) {
	userDB := &models2.User{}
	salt := ""
	qs := "select username, password, firstname, lastname, salt from users where username=$1 and password=$2"
	row := repo.service.QueryRow(context.Background(), qs, user.Username, user.Password)
	err := row.Scan(&userDB.Username, &userDB.Password, &userDB.Firstname, &userDB.Lastname, &salt)
	if err != nil {
		log.Println(err)
		return nil, models2.NOT_FOUND
	}
	if userDB.Password != repo.encr.EncryptString(user.Password, salt) {
		log.Println("mismatch passwords")
		return nil, models2.NOT_FOUND
	}
	userDB.Password = ""
	return userDB, models2.OK
}
