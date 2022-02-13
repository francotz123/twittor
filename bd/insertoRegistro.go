package bd

import (
	"context"
	"time"

	models "github.com/francotz123/twittor/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* InsertoRegistro es una funcion que se utilzia para registrar el usuario en la DB */
func InsertoRegistro(user models.Usuario) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("usuarios")

	user.Password, _ = EncriptarPassword(user.Password)

	result, err := collection.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
