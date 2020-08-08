package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/models"

	"go.mongodb.org/mongo-driver/bson"
)

// InsertoTweet funcion para guardar tweet de usuario
func InsertoTweet(tweet models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	registro := bson.M{
		"userid":  tweet.UserID,
		"mensaje": tweet.Mensaje,
		"fecha":   tweet.Fecha,
	}

	resultado, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}
	objID, _ := resultado.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil
}
