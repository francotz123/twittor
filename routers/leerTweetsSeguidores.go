package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	bd "github.com/francotz123/twittor/bd"
)

func LeerTweetsSeguidores(w http.ResponseWriter, r *http.Request) {
	paginaString := r.URL.Query().Get("pagina")

	if len(paginaString) < 1 {
		http.Error(w, "Debe pasar la pagina ", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(paginaString)

	if err != nil {
		http.Error(w, "Debe pasar la pagina como entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, status := bd.LeoTweetsSeguidores(IDUsuario, pagina)

	if !status {
		http.Error(w, "Hubo un error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
