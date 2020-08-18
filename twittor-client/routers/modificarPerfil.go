package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/bd"
	"github.com/Raylynd6299/Go-TwitterClone-Course.git/models"
)

// ModificarPerfil esta funcion modifica el registro en la BD
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario
	var status bool

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Datios Incorrectos "+err.Error(), 400)
		return
	}

	status, err = bd.ModificoRegistro(user, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el Registro. Reintente "+err.Error(), 400)
		return
	}
	if status != true {
		http.Error(w, "No de ha logrado modificar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
