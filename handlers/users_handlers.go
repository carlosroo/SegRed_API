package handlers

import (
    "fmt"
    "net/http"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"SEGRED_API/models"

)

// func login (user models.User){
// 	var jwtKey = []byte(secret_key)

// 	UsersDB = 

// 	expectedPassword, ok :=  

// }


func CreateUser (w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	var ruta string
	
	//Leer la peticion
	reqBody,  err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Inserte un usuario válido\n Error: %v", err)
		return
	}
	//Decodifico el json
	json.Unmarshal(reqBody, &newUser)
	if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Error al decodificar el JSON\n Error: %v", err)
        return
	}
	// Agrego el usuario a la bbdd
	if  err := CargarUsuarios(&UsersDB); err != nil {
		fmt.Println("Error al cargar la base de datos de usuarios:", err)
	}
	AgregarUsuario(&UsersDB, newUser.Name, newUser.Password)

	//Genero el token de usuario
	/*not implemented */

	//creo un nuevo directorio con el nombre del usuario
	ruta = filepath.Join(".", bbdd, newUser.Name)
	err2 := os.MkdirAll(ruta, 0755)
	w.Header().Set("Content-Type", "application/json")
	if err2!= nil {
		w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Inserte un usuario válido\n Error: %v", err2)
    } else {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Creado nuevo directorio de usuario: %v\n", newUser.Name)
	}
}

func GetVersion(w http.ResponseWriter, r *http.Request) { //manda un json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(version)
}

func IndexRoute(w http.ResponseWriter, r *http.Request){
	// w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Bienvenido lokete")
}