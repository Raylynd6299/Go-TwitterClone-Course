package routers

import (
	"encoding/json"
	"net/http"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/bd"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/models"
)

//Registro Esta funcion crea el registro en la BD de usuario
func Registro(w http.ResponseWriter, r *http.Request) {

	// variable para almacenar el usuario
	var user models.Usuario
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error en los Datos recibidos "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email de usuario es requerido ", 400)
		return
	}
	if len(user.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de al menos 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(user.Email)

	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(user)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
