package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/bd"
)

// LeoTweets funcion para gestionar los tweets y mandarlos a la pagina
func LeoTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe envaira el parametro pagina con un valor mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)
	respuesta, correcto := bd.LeoTweets(ID, pag)
	if correcto == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}
	w.Header().Set("Context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
