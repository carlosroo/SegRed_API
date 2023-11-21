package handlers

import {
	"fmt"
    "net/http"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"

	"SEGRED_API/models"
}

// Implementa /<string:username>/<string:doc_id>
func getFileContent(w http.ResponseWriter, r *http.Request){

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Token de autorización ausente")
		return
	}

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(authHeader, claims,
		func(t *jwt.Token) (interface{}, error){
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}



}

//Comprueba si hay token y si es valido

func validateToken(w *http.ResponseWriter) error {

}