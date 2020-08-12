package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN almacena la coneccion con la base de Datos
var MongoCN = ConectarBD()

// clienteOptions tiene tal cual a que cliente se conectara
var clienteOptions = options.Client().ApplyURI("mongodb+srv://ray:4802250Ray@twittor.qtmw7.mongodb.net/twittor?retryWrites=true&w=majority")

// ConectarBD es la funcion para optener la coneccion con la base de Datos de forma correcta
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clienteOptions)
	if err != nil {
		log.Fatal("Error conectando con mongo", err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Error en la BD de mongo", err.Error())
		return client
	}
	log.Println("Coneccion Exitosa a la BD")
	return client
}

// ChequeoConnection Comprueba la coneccion con la base de Datos
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
