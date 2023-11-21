package handlers

import (
    "fmt"
    "net/http"
	"encoding/json"
)

func GetVersion(w http.ResponseWriter, r *http.Request) { //manda un json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(version)
}

func IndexRoute(w http.ResponseWriter, r *http.Request){
	// w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Bienvenido lokete")
}