package middleware

import (
	"net/http"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/bd"
)

// ChequeoBD esta funcion es un paso intermedio entre la accion real y la web,
// para comprobar que la BD siga activa
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
