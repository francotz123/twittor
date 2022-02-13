package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb://ftelliz:1540257070@cluster0-shard-00-00.t6b0y.mongodb.net:27017,cluster0-shard-00-01.t6b0y.mongodb.net:27017,cluster0-shard-00-02.t6b0y.mongodb.net:27017/twittor?ssl=true&replicaSet=atlas-12rv9z-shard-0&authSource=admin&retryWrites=true&w=majority")

/* ConectarDB Es la funcion que permite conectar a la base de datos */
func ConectarDB() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("EROOR:" + err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexion existosa con la DB")
	return client
}

/* ChequeConnection Es la funcion que permite corroborar la conexion con un ping */
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
