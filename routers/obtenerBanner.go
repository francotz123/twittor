package routers

import (
	"io"
	"net/http"
	"os"

	bd "github.com/francotz123/twittor/bd"
)

/* ObtenerBanner devuelve el banner de un usuario */
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {

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

	openFile, err := os.Open("uploads/banners/" + perfil.Bannner)

	if err != nil {
		http.Error(w, "Error al buscar el banner del usuario", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, openFile)

	if err != nil {
		http.Error(w, "Error al devolver el banner del usuario", http.StatusBadRequest)
		return
	}

}
