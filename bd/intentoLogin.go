package bd

import (
	models "github.com/francotz123/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/* IntengoLogin realiza el chequeo de login a la DB*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	userModel, encontrado, _ := ChequeoYaExisteUsuario(email)
	if !encontrado {
		return userModel, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(userModel.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return userModel, false
	}

	return userModel, true
}
