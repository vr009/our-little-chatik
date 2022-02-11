package models

type User struct {
	Uuid      string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
