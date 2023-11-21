package main

import (
	"net/http"
	"log"

	"SEGRED_API/handlers"
	
	"github.com/gorilla/mux"
)


func main(){
	router := mux.NewRouter().StrictSlash(true)
	//rutas con la funcion que van a ejecutar cuando sean llamadas
	router.HandleFunc("/", handlers.IndexRoute).Methods("GET")
	router.HandleFunc("/version", handlers.GetVersion).Methods("GET")
	router.HandleFunc("/signup", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	log.Println("Servidor corriendo en el puerto 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
