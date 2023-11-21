package handlers

import (
    "fmt"
    "net/http"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"

	"SEGRED_API/models"

)
/*
*
* Métodos que implementan /login
*
*/
func identifyUser (user models.User) (string, error) {
	var jwtKey = []byte(secret_key)

	if  err := CargarUsuarios(); err != nil {
		fmt.Println("Error al cargar la base de datos de usuarios:", err)
		return tokenErroneo, err
	}
	user_bbdd, err := searchUserByName(user.Name)
	if err != nil {
		return tokenErroneo, err
	}
	if err := verifyPassword(user_bbdd, user.Password); err != nil {
		return tokenErroneo, err
	}
	expirationTime := time.Now().Add(time.Minute *5)

	claims := &models.Claims{
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString()

}

//buscamos al usuario en la base de datos
func searchUserByName(username string) (*models.User, error) {
	for _, usuario := range usersDB.Users {
		if usuario.Name == username {
			return &usuario, nil
		}
	}
	return nil, fmt.Errorf("Usuario no encontrado: %s", username)
}
//comparamos el hash de la contraseña recibida con el guaradado 
func verifyPassword(user *models.User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err
}

/*
*
* Métodos que implementan /signup
*
*/
func CreateUser (w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	
	//Leer la peticion
	reqBody,  err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error en la peticion\n Error: %v", err)
		return
	}
	//Decodifico el json
	err = json.Unmarshal(reqBody, &newUser)
	if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Error al decodificar el JSON\n Error: %v", err)
        return
	}
	//Incluyo el nuevo usuario en la base de datos
	err = addUserdb(newUser)
	if err!= nil {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Error al crear el usuario en la base de datos\n Error: %v", err)
        return
    }

	//Genero el token de usuario
	/*not implemented */

	//Creo un nuevo directorio para el usuario
	err = newDirectory(newUser.Name)
	w.Header().Set("Content-Type", "application/json")
	if err!= nil {
		w.WriteHeader(http.StatusBadRequest)
        fmt.Fprintf(w, "Error en la creacion del directorio\n Error: %v", err)
    } else {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Creado nuevo directorio de usuario: %v\n", newUser.Name)
	}
}
// Crea el directorio con un nombre del nuevo usuario
func newDirectory(dir string) error {
	ruta := filepath.Join(".", dir_usuarios, dir)
	err := os.MkdirAll(ruta, 0755)
	if err!= nil {
		return err
    } else {
		return nil
	}
}
// Agrego el nuevo usuario a la base de datos
func addUserdb (newUser models.User) error {
	if  err := CargarUsuarios(); err != nil {
		fmt.Println("Error al cargar la base de datos de usuarios:", err)
		return err
	}
	err1 := AddUser(newUser.Name, newUser.Password)
	if err1 != nil {
		return err1
    } else {
		return nil
	}
}
	