package repo

import (
	models2 "auth/internal/models"
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"strings"
)

type PGRepo struct {
	service *pgxpool.Pool
}

func NewPGRepo(service *pgxpool.Pool) *PGRepo {
	return &PGRepo{
		service: service,
	}
}

func (repo *PGRepo) CreateUser(user *models2.User) models2.ErrorCode {
	if repo.service != nil {
		uuidWithHyphen := uuid.New()
		uuidstr := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
		log.Println("new")
		str := "insert into users values ($1 ,$2, $3, $4, $5)"

		if _, err := repo.service.Exec(context.Background(),
			str, uuidstr, user.Username, user.Password, user.Firstname, user.Lastname); err != nil {
			log.Println(err)
			return models2.EXISTS
		}
	}
	return models2.OK
}

func (repo *PGRepo) GetUser(user *models2.User) (*models2.User, models2.ErrorCode) {
	userDB := &models2.User{}
	qs := "select username, password, firstname, lastname from users where username=$1 and password=$2"
	row := repo.service.QueryRow(context.Background(), qs, user.Username, user.Password)
	err := row.Scan(&userDB.Username, &userDB.Password, &userDB.Firstname, &userDB.Lastname)
	if err != nil {
		log.Println(err)
		return nil, models2.NOT_FOUND
	}
	return userDB, models2.OK
}
