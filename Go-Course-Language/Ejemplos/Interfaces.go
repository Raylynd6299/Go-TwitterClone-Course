package main

import "fmt"

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

/*Estructura Hombre*/

type hombre struct {
	edad       int
	altura     float64
	peso       float64
	respirando bool
	pensando   bool
	comiendo   bool
}

type mujer struct {
	edad       int
	altura     float64
	peso       float64
	respirando bool
	pensando   bool
	comiendo   bool
}

func (h *hombre) respirar()    { h.respirando = true }
func (h *hombre) comer()       { h.comiendo = true }
func (h *hombre) pensar()      { h.pensando = true }
func (h *hombre) sexo() string { return "Masculino" }

func (m *mujer) respirar()    { m.respirando = true }
func (m *mujer) comer()       { m.comiendo = true }
func (m *mujer) pensar()      { m.pensando = true }
func (m *mujer) sexo() string { return "Femenino" }

func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Println("Soy de sexo ", hu.sexo(), "y ya estoy respirando")
}

func main() {
	Ray := new(hombre)
	HumanosRespirando(Ray)

	Dani := new(mujer)
	HumanosRespirando(Dani)
}
