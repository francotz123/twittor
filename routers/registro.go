package routers

import (
	"encoding/json"
	"net/http"

	bd "github.com/francotz123/twittor/bd"

	models "github.com/francotz123/twittor/models"
)

/* Registro es la funcion para crear en la DB el registro de usuario */
func Registro(w http.ResponseWriter, r *http.Request) {
	var modelUser models.Usuario
	err := json.NewDecoder(r.Body).Decode(&modelUser)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(modelUser.Email) == 0 {
		http.Error(w, "Error: El email es requerido ", 400)
		return
	}

	if len(modelUser.Password) < 6 {
		http.Error(w, "Error: Debe especificar una contraseña de al menos 6 caracteres  ", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(modelUser.Email)

	if encontrado {
		http.Error(w, "Ya existe un usuario registrado con ese email.", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(modelUser)

	if err != nil {
		http.Error(w, "Error: Hubo un error al registrar usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logró insertar el registro del usuario ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
