package routers

import (
	"net/http"

	bd "github.com/francotz123/twittor/bd"
	models "github.com/francotz123/twittor/models"
)

func AltarRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe pasar el ID ", 400)
		return
	}

	var modelRelacion models.Relacion

	modelRelacion.UsuarioID = IDUsuario
	modelRelacion.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(modelRelacion)

	if err != nil {
		http.Error(w, "Error: Hubo un error al insertar la relacion "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logró insertar la relación ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
