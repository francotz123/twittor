package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	models "github.com/francotz123/twittor/models"
)

/* GeneroJWT genera el encriptado con JWT */
func GeneroJWT(userModel models.Usuario) (string, error) {

	miClave := []byte("FrancoTellizTwittor41981")

	payload := jwt.MapClaims{
		"email":            userModel.Email,
		"nombre":           userModel.Nombre,
		"apellidos":        userModel.Apellidos,
		"fecha_nacimiento": userModel.FechaNacimiento,
		"biografia":        userModel.Biografia,
		"ubicacion":        userModel.Ubicacion,
		"sitioweb":         userModel.Ubicacion,
		"_id":              userModel.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}
