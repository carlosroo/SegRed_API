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

	return ioutil.WriteFile(bbdd, usuariosCifrados, 0644)
}

func NuevoUsuario(username, password string) models.User {
	return models.User{
		Name: username,
		Password: password,
	}
}

func CargarUsuarios() error{
	fileContent, err := ioutil.ReadFile(bbdd)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileContent, &usersDB.Users)
}

func AddUser(nombre, contraseña string) error {
	// Verifica si el usuario ya existe
	for _, usuario := range usersDB.Users {
		if usuario.Name == nombre {
			return fmt.Errorf("el usuario %s ya existe", nombre)
		}
	}
	newUser := NuevoUsuario(nombre, cifrarContraseña(contraseña))

	usersDB.Users = append(usersDB.Users, newUser)
	
	if err:= GuardarUsuarios(&usersDB); err != nil {
		return err
	}

	return nil
}