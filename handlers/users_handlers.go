package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"SEGRED_API/models"
)

//funcion que implementa /login
func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error en la peticion\n Error: %v", err)
		return
	}
	//Decodifico el json
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error al decodificar el JSON\n Error: %v", err)
		return
	}
	err = identifyUser(&user, w)
	if err != nil {
		fmt.Fprintf(w, "Error en la autenticacion\n Error: %v", err)
		return
	}
	token, err := generateToken(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error en la genenacion del token: %v", err)
		return
	}
	//Inicio de sesion con cookie para no mandar el token en cada request
	// http.SetCookie(w,
	// 	&http.Cookie{
	// 		Name:    "token",
	// 		Value:   token,
	// 		Expires: expirationTime,
	// 		HttpOnly: true,
	// 	})

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"access_token": "%s"}`, token)
}

// verificar contraseña de usuario
func identifyUser(user *models.User, w http.ResponseWriter) error {

	if err := CargarUsuarios(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return fmt.Errorf("error al leer la base de datos de usuarios: %v", err)
	}
	user_bbdd, err := searchUserByName(user.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return fmt.Errorf("usuario no encontrado: %v", user.Username)
	}
	if err := verifyPassword(user_bbdd, user.Password); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return fmt.Errorf("contraseña incorrecta")
	}

	return nil
}

//genera y devuleve token de usuario
func generateToken(name string) (string, error) {
	expirationTime := time.Now().Add(time.Minute * token_expiration_time)

	claims := &models.Claims{
		Username: name,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

//genera y devuelve un valor aleatorio para usarlo como salt

//buscamos al usuario en la base de datos
func searchUserByName(username string) (*models.User, error) {
	for _, usuario := range usersDB.Users {
		if usuario.Username == username {
			return &usuario, nil
		}
	}
	return nil, fmt.Errorf("usuario no encontrado: %v", username)
}

//comparamos el hash de la contraseña recibida con el guaradado
func verifyPassword(user *models.User, password string) error {
	combined := append([]byte(password), []byte(user.Salt)...)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(combined))
	return err
}

//Metodo que implementa /signup
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	//Leer la peticion
	reqBody, err := ioutil.ReadAll(r.Body)
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
	//Genero salt
	newUser.Salt, err = generateSalt()
	if err != nil {
		w.WriteHeader((http.StatusInternalServerError))
		fmt.Fprintf(w, "error creando el usuario: %v", err)
	}
	//Incluyo el nuevo usuario en la base de datos
	err = addUserdb(newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error al crear el usuario en la base de datos\n Error: %v", err)
		return
	}

	//Genero el token de usuario
	token, err := generateToken(newUser.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error en la generacion del token de usuario\n Error: %v", err)
		return
	}

	//Creo un nuevo directorio para el usuario
	err = newDirectory(newUser.Username)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error en la creacion del directorio\n Error: %v", err)
		return
	} else {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, `{"access_token": "%v"}`, token)
		return
	}
}

// Crea el directorio con un nombre del nuevo usuario
func newDirectory(dir string) error {
	ruta := filepath.Join(".", dir_usuarios, dir)
	err := os.MkdirAll(ruta, 0755)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// Agrego el nuevo usuario a la base de datos
func addUserdb(newUser models.User) error {
	if err := CargarUsuarios(); err != nil {
		return fmt.Errorf("error al cargar la base de datos")
	}
	if newUser.Username == "" || newUser.Password == "" {
		return fmt.Errorf("usuario o contraseña vacío")
	}
	err1 := AddUser(newUser.Username, newUser.Password, newUser.Salt)
	if err1 != nil {
		return err1
	} else {
		return nil
	}
}
