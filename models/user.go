package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
	Salt     string `json:"salt"`
}

type UsersDB struct {
	Users []User
}
