package main

import (
	"log"
	"net/http"

	"SEGRED_API/handlers"

	"github.com/gorilla/mux"
)

func main() {
	certFile := "cert.pem"
	keyFile := "key_unencrypted.pem"
	router := mux.NewRouter().StrictSlash(true)

	err := handlers.InitializeDatabase()
	if err != nil {
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}

	RegisterRoutes(router)

	log.Println("Servidor corriendo en el puerto 5000")
	//log.Fatal(http.ListenAndServe(":5000", router))
	log.Fatal(http.ListenAndServeTLS(":5000", certFile, keyFile, router))
}

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", handlers.IndexRoute).Methods("GET")
	router.HandleFunc("/version", handlers.GetVersion).Methods("GET")
	router.HandleFunc("/signup", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/{username}/_all_docs", handlers.GetAllFiles).Methods("GET")
	router.HandleFunc("/{username}/{doc_id}", handlers.HandleFileOperations)
}
