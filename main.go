package main

import (
	"log"
	"net/http"

	"SEGRED_API/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handlers.IndexRoute).Methods("GET")
	router.HandleFunc("/version", handlers.GetVersion).Methods("GET")
	router.HandleFunc("/signup", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/{username}/{doc_id}", handlers.HandleFileOperations)

	log.Println("Servidor corriendo en el puerto 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
