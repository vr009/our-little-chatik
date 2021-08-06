package models

import "time"

type User struct {
	uuid     string
	UserName string `json:"UserName"`
	Password string `json:"Password"`
}

type message struct {
	Body      string
	Direction User
	Sender    User
	Sent      time.Time
}
