package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/bd"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/models"
)

// GraboTweet permite grabar el tweet en la BD
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Error Guardano el tweet "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ah guardado el tweet", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
