package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"

	"SEGRED_API/models"
)

//Crea y devuelve el hash de una cadena
func cifrarContraseña(password, salt string) string {
	combined := append([]byte(password), []byte(salt)...)
	hash, _ := bcrypt.GenerateFromPassword(combined, bcrypt.DefaultCost)
	return string(hash)
}

func generateSalt() (string, error) {
	saltBytes := make([]byte, 16)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(saltBytes), nil
}

//Vuelca los usuarios en la base de datos
func GuardarUsuarios(db *models.UsersDB) error {
	usuariosCifrados, err := json.Marshal(db.Users)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(bbdd, usuariosCifrados, 0644)
}

//Crea una nueva estructura User dados sus parametros
func NuevoUsuario(username, password, salt string) models.User {
	return models.User{
		Name:     username,
		Password: password,
		Salt:     salt,
	}
}

//Carga los usuarios de la base de datos a usersDB.Users
func CargarUsuarios() error {

	fileContent, err := ioutil.ReadFile(bbdd)
	if err != nil {
		return fmt.Errorf("error al leer la base de datos")
	}
	if len(fileContent) == 0 {
		return nil
	}
	return json.Unmarshal(fileContent, &usersDB.Users)
}

//Añade usuario a la base de datos
func AddUser(nombre, contraseña, salt string) error {
	// Verifica si el usuario ya existe
	for _, usuario := range usersDB.Users {
		if usuario.Name == nombre {
			return fmt.Errorf("el usuario %s ya existe", nombre)
		}
	}
	newUser := NuevoUsuario(nombre, cifrarContraseña(contraseña, salt), salt)

	usersDB.Users = append(usersDB.Users, newUser)

	if err := GuardarUsuarios(&usersDB); err != nil {
		return err
	}

	return nil
}

//Crear fichero de base de datos si no existe
func InitializeDatabase() error {

	_, err := os.Stat(bbdd)
	if os.IsNotExist(err) {
		// El archivo no existe, crearlo
		file, err := os.Create(bbdd)
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}
