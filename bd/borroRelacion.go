package bd

import (
	"context"
	"time"

	models "github.com/francotz123/twittor/models"
)

func BorroRelacion(modelRelacion models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("relacion")

	_, err := collection.DeleteOne(ctx, modelRelacion)

	if err != nil {
		return false, err
	}

	return true, nil
}
