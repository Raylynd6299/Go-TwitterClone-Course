package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/models"
)

// ChequeoYaExisteUsuario recibe un email de parametro y checa si ya esta en la BD
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()

	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
