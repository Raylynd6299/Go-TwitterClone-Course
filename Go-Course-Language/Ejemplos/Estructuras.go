package main

import (
	f "fmt"
	"time"
)

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}
type Admin struct {
	Usuario
	Permisos string
}

func (us *Usuario) CrearUsuario(id int, nombre string, fecha time.Time, status bool) {
	us.Id, us.Nombre = id, nombre
	us.FechaAlta, us.Status = fecha, status
}

func main() {
	User := new(Admin) // La palabra new crea un apuntador al tipo de dato entre parentesis

	User.CrearUsuario(1, "Raymudno Pulido", time.Now(), true)
	User.Permisos = "root"

	f.Println(User)

}
