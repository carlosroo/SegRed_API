package handlers

import (
	"fmt"
    "net/http"
	"path/filepath"
	"io/ioutil"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"

	"SEGRED_API/models"

)

// Implementa GET /<string:username>/<string:doc_id>
func GetFileContent(w http.ResponseWriter, r *http.Request){

	err:= validateToken(r.Header.Get("Authorization"))
	if err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Error en el token\n Error: %v", err)
		return 
	}

	vars:= mux.Vars(r)
	username:= vars["username"]

	docID := vars["doc_id"]
	if !strings.HasSuffix(docID, ".json") {
		docID += ".json"
	}

	filePath := filepath.Join(".", dir_usuarios, username, docID)
	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error al leer el archivo: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(fileContent)

}
// Implementa PUT /<string:username>/<string:doc_id>
func updateFileContent(w http.ResponseWriter, r *http.Request){
	err:= validateToken(r.Header.Get("Authorization"))
	if err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Error en el token\n Error: %v", err)
		return 
	}
	fmt.Fprintf(w, "Token validado con exito")
}
// Implementa DELETE /<string:username>/<string:doc_id>
func deleteFile(w http.ResponseWriter, r *http.Request){
	err:= validateToken(r.Header.Get("Authorization"))
	if err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Error en el token\n Error: %v", err)
		return 
	}
	fmt.Fprintf(w, "Token validado con exito")
}
// Implementa POST /<string:username>/<string:doc_id>
func uploadFile(w http.ResponseWriter, r *http.Request){
	err:= validateToken(r.Header.Get("Authorization"))
	if err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Error en el token\n Error: %v", err)
		return 
	}
	fmt.Fprintf(w, "Token validado con exito")
}
// Implementa GET /<string:username>/_all_docs
func getAllFiles(w http.ResponseWriter, r *http.Request){
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