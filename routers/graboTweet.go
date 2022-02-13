package routers

import (
	"encoding/json"
	"net/http"
	"time"

	bd "github.com/francotz123/twittor/bd"
	models "github.com/francotz123/twittor/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	if err != nil {
		http.Error(w, "Error en el mensaje recibido "+err.Error(), 400)
		return
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar guardar el Tweet. Reintente nuevamente."+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado guardar el Tweet correctamente", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
