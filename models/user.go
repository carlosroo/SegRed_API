package models

type User struct {
	Nombre      string `json:"nombre"`
	Contrasena  string `json:"contrasena"`
}

type User []users

