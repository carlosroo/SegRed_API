package handlers

import "SEGRED_API/models"

const (
	version = 0
	dir_usuarios = "data"
	secret_key = "secret_key"
	token_expiration_time = 5
	bbdd = "usuarios.db"
	tokenErroneo = "XXXXXXXXXXXX"
)

var usersDB models.UsersDB