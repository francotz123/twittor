package bd

import (
	"context"
	"log"
	"time"

	models "github.com/francotz123/twittor/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* ModificoRegistro permite modificar un registro usuario de la db  */
func ModificoRegistro(userModel models.Usuario, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	collection := db.Collection("usuarios")

	registro := make(map[string]interface{})

	if len(userModel.Nombre) > 0 {
		registro["nombre"] = userModel.Nombre
	}
	if len(userModel.Apellidos) > 0 {
		registro["apellidos"] = userModel.Apellidos
	}

	registro["fechaNacimiento"] = userModel.FechaNacimiento

	if len(userModel.Avatar) > 0 {
		registro["avatar"] = userModel.Avatar
	}
	if len(userModel.Bannner) > 0 {
		registro["banner"] = userModel.Bannner
	}
	if len(userModel.Biografia) > 0 {
		registro["biografia"] = userModel.Biografia
	}
	if len(userModel.Ubicacion) > 0 {
		registro["ubicacion"] = userModel.Ubicacion
	}
	if len(userModel.SitioWeb) > 0 {
		registro["sitioWeb"] = userModel.SitioWeb
	}

	updtString := bson.M{
		"$set": registro,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	condicion := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := collection.UpdateOne(ctx, condicion, updtString)

	if err != nil {
		log.Println("Error al acutalizar usuario" + err.Error())
		return false, err
	}

	return true, nil
}
