package models

import "time"

type Message struct {
	Body      string    `json:""`
	Direction string    `json:""` //uuid
	Sender    string    `json:""` //uuid
	Sent      time.Time `json:""`
}
