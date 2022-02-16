package routers

import (
	"fmt"
	"net/http"
)

func Bienvenido(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenid@ a la API de Twittor :)")
}
