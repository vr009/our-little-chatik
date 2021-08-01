package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"our-little-chatik/internal/models"
)

type PGRepo struct {
	service    *sql.DB
	Db_name    string
	Table_name string
}

func StartInit() {

}

func (repo *PGRepo) initDB() error {
	connStr := "user=postgres password=admin sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	repo.service = db

	return nil
}

func (repo *PGRepo) CreateUser(user models.User) error {
	if repo.service != nil {
		if _, err := repo.service.Exec("insert into %s values (%s, %s);", repo.Table_name, user.UserName, user.Password); err != nil {
			return err
		}
	}
	return nil
}

func (repo *PGRepo) GetUser(user models.User) (string, error) {

	out := ""
	if repo.service != nil {
		res, err := repo.service.Query(fmt.Sprintf("select user_id from %s", repo.Db_name))
		if err != nil {
			return "", err
		}
		res.Scan(&out)
		return out, nil
	}
	return "", errors.New("No connection")
}

func (repo *PGRepo) Close() {
	repo.service.Close()
}
