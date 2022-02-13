package main

import (
	"log"

	bd "github.com/francotz123/twittor/bd"
	handlers "github.com/francotz123/twittor/handlers"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin Conexion a la Base de Datos")
		return
	}

	handlers.Manejadores()
}
