package models

import "time"

type User struct {
	//Id       string `json:"uuid"`
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}

type message struct {
	Body      string
	Direction User
	Sender    User
	Sent      time.Time
}
