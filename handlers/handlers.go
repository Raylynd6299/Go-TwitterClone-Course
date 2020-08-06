package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores Configuro el puerto el handler y pongo a escuchar al servidor
func Manejadores() {
	router := mux.NewRouter()

	//Puerto a manejar
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	//cors permisos para la api para ser accesible desde cualquier lugar
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
