package routers

import (
	"net/http"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/bd"
	"github.com/Raylynd6299/Go-TwitterClone-Course.git/models"
)

// AltaRelacion realiza el registro de la relacion entre usuario
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro id", http.StatusBadRequest)
		return
	}
	var rel models.Relacion

	rel.UsuarioID = IDUsuario
	rel.UsuarioRealcionID = ID
	status, err := bd.InsertoRelacion(rel)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar la realcion"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
