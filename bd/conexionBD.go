package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://ftelliz:1540257070@cluster0.t6b0y.mongodb.net/twittor?retryWrites=true&w=majority")

/* ConectarDB Es la funcion que me permite conectar a la base de datos */
func ConectarDB() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("EROOR:" + err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}

	log.Println("Conexion existosa con la DB")
	return client
}

/* ChequeConnection Es la funcion que me permite corroborar la conexion con un ping */
func ChequeConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
