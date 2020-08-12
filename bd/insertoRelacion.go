package bd

import (
	"context"
	"time"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/models"
)

// InsertoRelacion crear relacion en bd
func InsertoRelacion(rel models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, rel)

	if err != nil {
		return false, err
	}
	return true, nil
}
