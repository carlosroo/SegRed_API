package models

import "github.com/dgrijalva/jwt-go"


type Claims struct {
    Username  string `json:"name"`
	jwt.StandardClaims 
}