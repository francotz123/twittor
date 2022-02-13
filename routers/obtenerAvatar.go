package routers

import (
	"io"
	"net/http"
	"os"

	bd "github.com/francotz123/twittor/bd"
)

/* ObtenerAvatar devuelve el avatar de un usuario */
func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe pasar un ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Error al buscar el perfil del usuario", http.StatusBadRequest)
		return
	}

	openFile, err := os.Open("uploads/avatars/" + perfil.Avatar)

	if err != nil {
		http.Error(w, "Error al buscar el avatar del usuario", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)

	if err != nil {
		http.Error(w, "Error al devolver el avatar del usuario", http.StatusBadRequest)
		return
	}

}
