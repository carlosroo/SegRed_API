package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil" // manejar entradas y salidas
	"os"
	"path/filepath"
	

	"github.com/gorilla/mux"
)

const version = 0
const bbdd = "data"

type User struct {
	Nombre      string `json:"nombre"`
	Contrasena  string `json:"contrasena"`
}


func createUser (w http.ResponseWriter, r *http.Request) { //solicita un json y crea una carpetilla
	var newUser User
	var ruta string
	var name string

	//leer la peticion y control de errores
	reqBody,  err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Inserte un usuario válido\n Error: %v", err)
		return
	}

	//descodifico el json
	json.Unmarshal(reqBody, &newUser)
	name = newUser.Nombre
	if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Error al decodificar el JSON\n Error: %v", err)
        return
	}

	//creo un nuevo directorio con el nombre del usuario
	ruta = filepath.Join(".", bbdd, name)
	err2 := os.MkdirAll(ruta, 0755)
	w.Header().Set("Content-Type", "application/json")
	if err2!= nil {
		w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Inserte un usuario válido\n Error: %v", err2)
    } else {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Creado nuevo directorio de usuario: %v\n", name)
	}
}

func getVersion(w http.ResponseWriter, r *http.Request) { //manda un json
	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(version)
}

func indexRoute(w http.ResponseWriter, r *http.Request){
	// w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Bienvenido lokete")
}

func main(){
	router := mux.NewRouter().StrictSlash(true)
	//rutas con la funcion que van a ejecutar cuando sean llamadas
	router.HandleFunc("/", indexRoute).Methods("GET")
	router.HandleFunc("/version", getVersion).Methods("GET")
	router.HandleFunc("/signup", createUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", router))
}
