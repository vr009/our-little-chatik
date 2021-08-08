package models

type User struct {
	uuid      string
	Firstname string `json:"Firstname,omitempty"`
	Lastname  string `json:"Lastname,omitempty"`
	Username  string `json:"Username"`
	Password  string `json:"Password"`
}
