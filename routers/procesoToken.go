package routers

import (
	"errors"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	bd "github.com/francotz123/twittor/bd"
	models "github.com/francotz123/twittor/models"
)

/* Email es el valor del email devuelto por el Token */
var Email string

/* IDUsuario es el valor del ID del usuario devuelto por el Token */
var IDUsuario string

/* ProcesoToken procesa roken para extraer sus valores */
func ProcesoToken(token string) (*models.Claim, bool, string, error) {
	miClave := []byte(os.Getenv("CLAVE_TOKEN"))
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err != nil {
		if !tkn.Valid {
			return claims, false, string(""), errors.New("token invalido")
		}
	}

	_, encontrado, ID := bd.ChequeoYaExisteUsuario(claims.Email)

	if encontrado {
		Email = claims.Email
		IDUsuario = ID
	}

	return claims, true, IDUsuario, nil
}
