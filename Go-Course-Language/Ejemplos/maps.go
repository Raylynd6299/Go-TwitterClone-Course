package main

import (
	"fmt"
)

func main() {
	//Crear mapa
	paises := make(map[string]string)

	fmt.Println(paises)

	paises["Mexico"] = "CDMX"
	fmt.Println(paises["Mexico"])

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}

	campeonato["Chivas"] = 55
	campeonato["Rayados"] = 40

	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Println(equipo, puntaje)
	}
}
