package middlew

import (
	"net/http"

	bd "github.com/francotz123/twittor/bd"
)

/* ChequeoBD es el middleware que permite conocer el estado de la DB */
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexión perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
