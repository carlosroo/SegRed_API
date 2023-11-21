package handlers

import (

	"encoding/json"
	"io/ioutil"
	"fmt"
	
	"golang.org/x/crypto/bcrypt"

	"SEGRED_API/models"
)

func cifrarContraseña(password string) string{
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}

func GuardarUsuarios(db *models.UsersDB) error {
	usuariosCifrados, err := json.Marshal(db.Users)
	if err != nil {
		return err
	}

	return ioutil.WriteFile("usuarios.db", usuariosCifrados, 0644)
}

func NuevoUsuario(username, password string) models.User {
	return models.User{
		Name: username,
		Password: password,
	}
}

func CargarUsuarios(db *models.UsersDB) error{
	usuariosCifrados, err := ioutil.ReadFile("./usuarios.db")
	if err != nil {
		return err
	}

	return json.Unmarshal(usuariosCifrados, &db.Users)
}

func AgregarUsuario(db *models.UsersDB, nombre, contraseña string) error {
	// Verifica si el usuario ya existe
	for _, usuario := range db.Users {
		if usuario.Name == nombre {
			return fmt.Errorf("el usuario %s ya existe", nombre)
		}
	}
	newUser := NuevoUsuario(nombre, cifrarContraseña(contraseña))
	//hash de la contraseña

	db.Users = append(db.Users, newUser)
	
	if err:= GuardarUsuarios(db); err != nil {
		return err
	}

	return nil
}