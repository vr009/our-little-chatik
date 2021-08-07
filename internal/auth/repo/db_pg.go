package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"our-little-chatik/internal/models"
	"strings"
)

type PGRepo struct {
	service    *sql.DB
	Db_name    string
	Table_name string
}

func NewPGRepo() *PGRepo {
	return &PGRepo{Db_name: "postgres", Table_name: "Users"}
}

func (repo *PGRepo) StartInit() error {
	create_query :=
		"create table Users (user_id uuid default uuid_generate_v4()," +
			" username varchar(50) primary key not null," +
			" password varchar(150)," +
			" firstname varchar(50)," +
			" lastname varchar(50));"

	if repo.service != nil {
		if _, err := repo.service.Exec(create_query); err != nil {
			return err
		}
	}
	return nil
}

func (repo *PGRepo) InitDB() error {
	connStr := "user=postgres password=admin dbname=postgres sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	repo.service = db

	return nil
}

func (repo *PGRepo) CreateUser(user models.User) (string, error) {
	if repo.service != nil {
		uuidWithHyphen := uuid.New()
		uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)

		str := fmt.Sprintf("insert into %s values ( '%s' ,'%s', '%s','%s', '%s');",
			repo.Table_name, uuid, user.Username, user.Password, user.Firstname, user.Lastname)

		if _, err := repo.service.Exec(str); err != nil {
			return uuid, err
		}
	}
	return "", nil
}

func (repo *PGRepo) GetUser(user models.User) (string, string, error) {

	if repo.service != nil {
		var pswd string
		str := fmt.Sprintf("select password from %s where username='%s';", repo.Table_name, user.Username)
		res := repo.service.QueryRow(str)
		err := res.Scan(&pswd)
		if err != nil {
			return "", "", err
		}

		var uuid string
		str2 := fmt.Sprintf("select user_id from %s where username='%s';", repo.Table_name, user.Username)
		res2 := repo.service.QueryRow(str2)
		err2 := res2.Scan(&uuid)
		if err2 != nil {
			return "", "", err
		}

		return uuid, pswd, nil
	}
	return "", "", errors.New("No connection")
}

func (repo *PGRepo) Close() {
	repo.service.Close()
}
