package routers

import (
	"encoding/json"
	"net/http"

	bd "github.com/francotz123/twittor/bd"
	models "github.com/francotz123/twittor/models"
)

func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var modelUser models.Usuario
	err := json.NewDecoder(r.Body).Decode(&modelUser)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(modelUser, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro. Reintente nuevamente."+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar el registro de usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
