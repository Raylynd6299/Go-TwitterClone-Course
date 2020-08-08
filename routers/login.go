package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/bd"
	"github.com/Raylynd6299/Go-TwitterClone-Course.git/jwt"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/models"
)

// Login realiza login
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var user models.Usuario

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "usuario y/o Contraseña invalidos "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email del usuario es requerido ", 400)
		return
	}
	documento, existe := bd.IntentoLogin(user.Email, user.Password)

	if existe == false {
		http.Error(w, "usuario y/o Contraseña invalidos ", 400)
		return
	}

	jwtKey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrio error generando el token "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	//Grabar una Cookie (no es usado en este curso)
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
