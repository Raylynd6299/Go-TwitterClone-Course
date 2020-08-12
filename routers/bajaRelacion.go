package routers

import (
	"net/http"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/bd"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/models"
)

// BajaRelacion gestiona la relacion entre usuarios
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var rel models.Relacion
	rel.UsuarioID = IDUsuario
	rel.UsuarioRealcionID = ID
	status, err := bd.BorroRelacion(rel)

	if err != nil {
		http.Error(w, "Ocurrio un error al eliminar la realcion"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
