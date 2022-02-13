package bd

import (
	"context"
	"time"

	models "github.com/francotz123/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoTweet permite guardar un tweet en la db */
func InsertoTweet(modelTweet models.GraboTweet) (string, bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("tweet")

	registro := bson.M{
		"userid":  modelTweet.UserID,
		"mensaje": modelTweet.Mensaje,
		"fecha":   modelTweet.Fecha,
	}

	result, err := collection.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.Hex(), true, nil
}
