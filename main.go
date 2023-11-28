package main

import (
	"log"
	"net/http"

	"SEGRED_API/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	err := handlers.InitializeDatabase()
	if err != nil {
		println("Error al inicializar la base de datos:", err)
		return
	}

	router.HandleFunc("/", handlers.IndexRoute).Methods("GET")
	router.HandleFunc("/version", handlers.GetVersion).Methods("GET")
	router.HandleFunc("/signup", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/{username}/_all_docs", handlers.GetAllFiles).Methods("GET")
	router.HandleFunc("/{username}/{doc_id}", handlers.HandleFileOperations)

	log.Println("Servidor corriendo en el puerto 5000")
	log.Fatal(http.ListenAndServe("127.0.0.1:5000", router))
}
