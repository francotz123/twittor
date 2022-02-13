package bd

import (
	"context"
	"time"

	models "github.com/francotz123/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* ChequeoYaExisteUsario es una funcion que verifica si ya existe un usuario registrado con ese email */
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := collection.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}

	return resultado, true, ID

}
