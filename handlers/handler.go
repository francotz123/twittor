package handlers

import (
	"log"
	"net/http"
	"os"

	middlew "github.com/francotz123/twittor/middlew"
	routers "github.com/francotz123/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/healthy", middlew.ChequeoBD(routers.Healthy)).Methods("GET")
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	err := http.ListenAndServe(":"+PORT, handler)
	if err != nil {
		log.Fatal(err.Error())
	}
}
