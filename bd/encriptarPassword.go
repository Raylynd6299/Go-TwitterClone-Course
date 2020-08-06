package bd

import (
	"golang.org/x/crypto/bcrypt"
)

// EncriptarPassword encripta la cadena con un costo de 8
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
