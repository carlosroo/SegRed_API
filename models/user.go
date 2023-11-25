package models

type User struct {
	Name     string `json:"name"`
	Password string `json:"password,omitempty"`
	Salt     string `json:"salt"`
}

type UsersDB struct {
	Users []User
}
