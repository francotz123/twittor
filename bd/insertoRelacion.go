package bd

import (
	"context"
	"log"
	"time"

	models "github.com/francotz123/twittor/models"
)

/* InsertoRelacion inserta en la db la relacion entre dos usuarios */
func InsertoRelacion(modelRelacion models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("relacion")

	_, err := collection.InsertOne(ctx, modelRelacion)

	if err != nil {
		log.Fatal("Hubo un error al insertar la relacion " + err.Error())
		return false, err
	}

	return true, nil

}
