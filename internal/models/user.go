package models

type User struct {
	Uuid      string `json:"-"`
	Firstname string `json:"Firstname,omitempty"`
	Lastname  string `json:"Lastname,omitempty"`
	Username  string `json:"Username"`
	Password  string `json:"Password"`
}
