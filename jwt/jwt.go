package jwt

import (
	"time"

	"github.com/Raylynd6299/Go-TwitterClone-Course.git/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GeneroJWT genera el encryptado con jwt
func GeneroJWT(user models.Usuario) (string, error) {
	miClave := []byte("RaymundoPulidoBejarano")

	payload := jwt.MapClaims{
		"email":           user.Email,
		"nombre":          user.Nombre,
		"apellidos":       user.Apellidos,
		"fechaNacimiento": user.FechaNacimiento,
		"biografia":       user.Biografia,
		"ubicacion":       user.Ubicacion,
		"sitioweb":        user.SitioWeb,
		"_id":             user.ID.Hex(),
		"exp":             time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
