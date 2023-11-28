package handlers

import (
	"fmt"
	"net/http"
)

//Implementa GET /version
func GetVersion(w http.ResponseWriter, r *http.Request) { //manda un json

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("%v", version)))
}

//Implementa GET /
func IndexRoute(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Welcome!")
}
