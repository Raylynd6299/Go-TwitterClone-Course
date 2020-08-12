package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/bd"
	"github.com/Raylynd6299/Go-TwitterClone-Course.git/models"
)

// ConsultaRelacion checa si hay relacion entre 2 usuarios
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var rel models.Relacion

	rel.UsuarioID = IDUsuario
	rel.UsuarioRealcionID = ID

	var resp models.RespuestaConsultaRelacion
	status, err := bd.ConsultoRelacion(rel)
	if err != nil || status != true {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
