package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/middleware"
	"github.com/Raylynd6299/Go-TwitterClone-Course.git/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Manejadores Configuro el puerto el handler y pongo a escuchar al servidor
func Manejadores() {
	//Manejador de Rutas
	router := mux.NewRouter()

	//Rutas
	router.HandleFunc("/registro", middleware.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleware.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middleware.ChequeoBD(middleware.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middleware.ChequeoBD(middleware.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middleware.ChequeoBD(middleware.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweets", middleware.ChequeoBD(middleware.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middleware.ChequeoBD(middleware.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middleware.ChequeoBD(middleware.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middleware.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")

	router.HandleFunc("/subirBanner", middleware.ChequeoBD(middleware.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middleware.ChequeoBD(routers.ObtenerBanner)).Methods("GET")

	//Puerto a manejar
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	//cors permisos para la api para ser accesible desde cualquier lugar
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
