package main

import (
	"log"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/bd"
	"github.com/Raylynd6299/Go-TwitterClone-Course.git/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
