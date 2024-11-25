package handlers

import "SEGRED_API/models"

const (
	version               = "2.0.0"
	dir_usuarios          = "data"
	secret_key            = "secret_key"
	token_expiration_time = 5
	bbdd                  = "usuarios.db"
	tokenErroneo          = ""
)

var jwtKey = []byte(secret_key)
var usersDB models.UsersDB
