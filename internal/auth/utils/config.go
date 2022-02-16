package utils

import (
	"fmt"
	"os"
)

func ConnStr() (string, error) {
	//errFunc := func(envVar string) (string, error) {
	//	return "", errors.Errorf("%s is not set in env vars", envVar)
	//}

	connstr := "port=%s host=%s password=%s user=%s dbname=%s"

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5433"
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "db-auth"
	}
	password := os.Getenv("DB_PSWD")
	if password == "" {
		password = "pswd"
	}
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "user"
	}
	name := os.Getenv("DB_NAME")
	if name == "" {
		name = "name"
	}
	return fmt.Sprintf(connstr, port, host, password, user, name), nil
}
