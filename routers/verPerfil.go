package routers

import (
	"encoding/json"
	"net/http"

	bd "github.com/francotz123/twittor/bd"
)

/* VerPefil premite extrae los valores del Perfil */
func VerPerfil(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el párametro ID ", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "No se encontró el Perfil"+err.Error(), 400)
		return
	}

	w.Header().Set("Context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
