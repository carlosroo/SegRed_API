package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"

	"SEGRED_API/models"
)

// GET /<string:username>/<string:doc_id>
func HandleFileOperations(w http.ResponseWriter, r *http.Request) {

	if err := handleToken(w, r); err != nil {
		return
	}

	vars := mux.Vars(r)
	username := vars["username"]
	docID := vars["doc_id"]

	if !strings.HasSuffix(docID, ".json") {
		docID += ".json"
	}

	filePath := filepath.Join(".", dir_usuarios, username, docID)

	switch r.Method {
	case http.MethodGet:
		getFileContent(w, filePath)
	case http.MethodPost:
		uploadFile(w, r, filePath)
	case http.MethodPut:
		updateFileContent(w, r, filePath)
	case http.MethodDelete:
		deleteFile(w, filePath)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "método no permitido: %s", r.Method)
	}

}

// POST /<string:username>/<string:doc_id> Crea un fichero json si no existe
func uploadFile(w http.ResponseWriter, r *http.Request, filePath string) {
	//Comprobar que el archivo existe
	_, err := os.Stat(filePath)
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "El archivo ya existe: %s", filePath)
		return
	} else if !os.IsNotExist(err) {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error al verificar la existencia del archivo: %v", err)
		return
	}
	// Subo el contenido en un nuevo archivo
	fileContent, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error al leer el cuerpo de la solicitud: %v", err)
		return
	}
	err = ioutil.WriteFile(filePath, fileContent, 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error al escribir en el archivo: %v", err)
		return
	}
	//Devulevo el tamaño del archivo
	fileStat, err := os.Stat(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error al obtener el tamaño del archivo: %v", err)
		return
	}

	response := fmt.Sprintf(`{"size": "%d"}`, fileStat.Size())
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, response)

}

// GET /<string:username>/<string:doc_id> Devuelve el contenido de un fichero json en una ruta
func getFileContent(w http.ResponseWriter, filePath string) {

	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error al leer el archivo: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(fileContent)

}

// DELETE /<string:username>/<string:doc_id> //Elimina un fichero json si existe
func deleteFile(w http.ResponseWriter, filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		//Compruebo si el archivo existe
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "El archivo no existe")
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error al borrar el archivo")
		return
	}
	w.WriteHeader(http.StatusOK)
}

// PUT /<string:username>/<string:doc_id>
func updateFileContent(w http.ResponseWriter, r *http.Request, filePath string) {
	fileContent, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error al leer el cuerpo de la solicitud: %v", err)
		return
	}
	//Verificar si el archivo existe
	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// El archivo no existe, responder con 404 (Not Found)
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "El archivo no existe: %s", filePath)
			return
		}

		// Otro tipo de error, manejar según sea necesario
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error al verificar la existencia del archivo: %v", err)
		return
	}
	//Sustituir el contenido del archivo
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error al abrir el archivo: %v", err)
		return
	}
	defer file.Close()

	_, err = file.Write(fileContent)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error al escribir en el archivo: %v", err)
		return
	}
	// Devulevo el tamaño del archivo
	fileStat, err := os.Stat(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error al obtener el tamaño del archivo: %v", err)
		return
	}

	response := fmt.Sprintf(`{"size": "%d"}`, fileStat.Size())
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, response)
}

// Implementa GET /<string:username>/_all_docs
func GetAllFiles(w http.ResponseWriter, r *http.Request) {
	if err := handleToken(w, r); err != nil {
		return
	}

	vars := mux.Vars(r)
	username := vars["username"]

	dirPath := filepath.Join(".", dir_usuarios, username)
	//leo el directorio de usuario
	userFiles, err := ioutil.ReadDir(dirPath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error al leer el directorio del usuario: %v", err)
		return
	}

	userData := make(map[string]interface{})

	for _, fileInfo := range userFiles {
		if fileInfo.IsDir() {
			continue
		}
		docID := strings.TrimSuffix(fileInfo.Name(), ".json")
		filePath := filepath.Join(dirPath, fileInfo.Name())
		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error al leer el archivo %v del usuario: %v", fileInfo.Name(), err)
			continue
		}

		var docContent interface{}
		err = json.Unmarshal(fileContent, &docContent)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error al decodificar el archivo %v del usuario: %v", fileInfo.Name(), err)
			continue
		}
		userData[docID] = docContent
	}
	response, err := json.Marshal(userData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error al generar la respuesta JSON: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

// Comprueba si hay token y si es valido
func validateToken(authHeader string) error {
	//authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		//w.WriteHeader(http.StatusUnauthorized)
		//fmt.Fprintf(w, "Token de autorización ausente")
		return fmt.Errorf("token de autorizacion ausente")
	}

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(authHeader, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			//w.WriteHeader(http.StatusUnauthorized)
			return fmt.Errorf("no autorizado")
		}
		//w.WriteHeader(http.StatusBadRequest)
		return fmt.Errorf("error en la solicitud")
	}

	if !tkn.Valid {
		//w.WriteHeader(http.StatusUnauthorized)
		return fmt.Errorf("token no válido")
	}
	return nil
}

// Valida el token y gestiona el error
func handleToken(w http.ResponseWriter, r *http.Request) error {
	headerAuth := r.Header.Get("Authorization")
	headerSplit := strings.Fields(headerAuth)
	token := headerSplit[1]

	err := validateToken(token)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Error en el token\n Error: %v", err)
		return err
	}
	return nil
}
