package routers

import (
	"encoding/json"
	"net/http"
	"time"

	bd "github.com/francotz123/twittor/bd"
	jwt "github.com/francotz123/twittor/jwt"
	models "github.com/francotz123/twittor/models"
)

/* Login realiza el login */
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")

	var modelUser models.Usuario
	err := json.NewDecoder(r.Body).Decode(&modelUser)

	if err != nil {
		http.Error(w, "Error al recibir el email y contraseña "+err.Error(), 400)
		return
	}

	if len(modelUser.Email) == 0 {
		http.Error(w, "El email es requerido ", 400)
		return
	}

	documentUser, existe := bd.IntentoLogin(modelUser.Email, modelUser.Password)

	if !existe {
		http.Error(w, "Usuario y/o contraseña inválidos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documentUser)

	if err != nil {
		http.Error(w, "Ocurrio un error al general el Token correspondiente"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 + time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
