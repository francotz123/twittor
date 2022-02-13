package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	bd "github.com/francotz123/twittor/bd"
)

/* LeoTweets devuelve los tweets*/
func LeoTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe pasar el ID ", 400)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe pasar el nro de página ", 400)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "La página no es un numero mayor a 0 ", 400)
		return
	}

	pag := int64(pagina)

	respuesta, correcto := bd.LeoTweets(ID, pag)

	if !correcto {
		http.Error(w, "Hubo un error al leer los tweets"+err.Error(), 400)
		return
	}

	w.Header().Set("Context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
