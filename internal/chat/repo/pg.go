package repo

import "database/sql"

type PGRepo struct {
	service    *sql.DB
	Db_name    string
	Table_name string
}

func NewPGRepo() *PGRepo {
	return &PGRepo{Db_name: "postgres", Table_name: "Users"}
}
