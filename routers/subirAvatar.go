package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	bd "github.com/francotz123/twittor/bd"
	models "github.com/francotz123/twittor/models"
)

/* SubirAvatar se actualiza el avatar del usuario */
func SubirAvatar(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")

	var extension = strings.Split(handler.Filename, ".")[1]

	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

	fle, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 666)

	if err != nil {
		http.Error(w, "Error al subir el avatar "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(fle, file)

	if err != nil {
		http.Error(w, "Error al copiar el avatar "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario

	var status bool

	usuario.Avatar = IDUsuario + "." + extension

	status, err = bd.ModificoRegistro(usuario, IDUsuario)

	if err != nil || !status {
		http.Error(w, "Error al guardar el avatar en la db "+err.Error(), 400)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
