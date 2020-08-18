package routers

import (
	"errors"
	"strings"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/bd"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Email valor de email usado en todos los Endpoints
var Email string

// IDUsuario es el ID devuelto del model, que se usara en todos los Endpoints
var IDUsuario string

// ProcesoToken proceso token para extraer sus valores
func ProcesoToken(token string) (*models.Claim, bool, string, error) {
	miClave := []byte("RaymundoPulidoBejarano")

	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, "", errors.New("formato de token invalido")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(toke *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, "", errors.New("formato de token invalido")
	}
	return claims, false, "", err
}
