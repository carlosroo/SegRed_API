package handlers

import (
	"fmt"
    "net/http"
	// "encoding/json"
	// "io/ioutil"
	// "os"
	// "path/filepath"
	// "time"

	// "golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"

	"SEGRED_API/models"
)

// Implementa /<string:username>/<string:doc_id>
func getFileContent(w http.ResponseWriter, r *http.Request){

	err:= validateToken(r.Header.Get("Authorization"))
	if err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Error en el token\n Error: %v", err)
		return 
	}
	fmt.Fprintf(w, "Token validado con exito")
}

//Comprueba si hay token y si es valido
func validateToken(authHeader string) error {
	//authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		//w.WriteHeader(http.StatusUnauthorized)
		//fmt.Fprintf(w, "Token de autorización ausente")
		return fmt.Errorf("Token de autorizacion ausente")
	}

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(authHeader, claims,
		func(t *jwt.Token) (interface{}, error){
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			//w.WriteHeader(http.StatusUnauthorized)
			return fmt.Errorf("No autorizado")
		}
		//w.WriteHeader(http.StatusBadRequest)
		return fmt.Errorf("Error en la solicitud")
	}

	if !tkn.Valid {
		//w.WriteHeader(http.StatusUnauthorized)
		return fmt.Errorf("Token no válido")
	}
	return nil
}