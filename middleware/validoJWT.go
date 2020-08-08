package middleware

import (
	"log"
	"net/http"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/routers"
)

// ValidoJWT permite validar el JWT que nos viene en la peticion
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el token ! "+err.Error(), http.StatusBadRequest)
			log.Println("Error aqui")
			return
		}
		next.ServeHTTP(w, r)
	}
}
