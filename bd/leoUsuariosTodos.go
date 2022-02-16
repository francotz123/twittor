package bd

import (
	"context"
	"fmt"
	"time"

	models "github.com/francotz123/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LeoUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("usuarios")

	var resultados []*models.Usuario

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSkip((page - 1) * 20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := collection.Find(ctx, query, opciones)

	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}

	var encontrado, incluir bool

	for cursor.Next(ctx) {
		var modelUser models.Usuario
		err := cursor.Decode(&modelUser)

		if err != nil {
			fmt.Println(err.Error())
			return resultados, false
		}
		var relacion models.Relacion

		relacion.UsuarioID = ID
		relacion.UsuarioRelacionID = modelUser.ID.Hex()

		incluir = false
		encontrado, err = ConsultoRelacion(relacion)

		if tipo == "new" && !encontrado {
			incluir = true
		}

		if tipo == "follow" && encontrado {
			incluir = true
		}

		if relacion.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir {
			modelUser.Password = ""
			modelUser.Biografia = ""
			modelUser.Bannner = ""
			modelUser.Email = ""
			modelUser.SitioWeb = ""
			modelUser.Ubicacion = ""

			resultados = append(resultados, &modelUser)
		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return resultados, false
	}
	cursor.Close(ctx)

	return resultados, true
}
