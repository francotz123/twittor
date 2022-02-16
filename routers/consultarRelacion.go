package routers

import (
	"encoding/json"
	"net/http"

	bd "github.com/francotz123/twittor/bd"
	models "github.com/francotz123/twittor/models"
)

func ConsultarRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe pasar el ID ", 400)
		return
	}

	var modelRelacion models.Relacion

	modelRelacion.UsuarioID = IDUsuario
	modelRelacion.UsuarioRelacionID = ID

	status, _ := bd.ConsultoRelacion(modelRelacion)

	resp := models.RespuestaConsultaRelacion{
		Status: status,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
