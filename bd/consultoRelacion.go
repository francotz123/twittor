package bd

import (
	"context"
	"fmt"
	"time"

	models "github.com/francotz123/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ConsultoRelacion(modelRelacion models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         modelRelacion.UsuarioID,
		"usuariorelacionid": modelRelacion.UsuarioRelacionID,
	}

	var resultado models.Relacion

	fmt.Print(resultado)

	err := collection.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		return false, err
	}

	return true, nil

}
