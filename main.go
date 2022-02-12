package main

import (
	"log"

	bd "github.com/francotz123/twittor/bd"
	handlers "github.com/francotz123/twittor/handlers"
)

func main() {
	if bd.ChequeConnection() == 0 {
		log.Fatal("Sin Conexion a la Base de Datos")
		return
	}

	handlers.Manejadores()
}
