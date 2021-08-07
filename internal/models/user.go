package models

import "time"

type User struct {
	uuid      string
	Firstname string `json:"Firstname,omitempty"`
	Lastname  string `json:"Lastname,omitempty"`
	Username  string `json:"Username"`
	Password  string `json:"Password"`
}

type Message struct {
	Body      string    `json:""`
	Direction User      `json:""`
	Sender    User      `json:""`
	Sent      time.Time `json:""`
}
