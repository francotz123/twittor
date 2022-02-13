package bd

import (
	"context"
	"log"
	"time"

	models "github.com/francotz123/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* BuscoPefil busca un perfil en la db */
func BuscoPerfil(ID string) (models.Usuario, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("usuarios")

	var perfil models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{"_id": objID}

	err := collection.FindOne(ctx, condicion).Decode(&perfil)
	perfil.Password = ""

	if err != nil {
		log.Println("Error al buscar usuario" + err.Error())
		return perfil, err
	}

	return perfil, nil

}
