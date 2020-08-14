package main

import (
	"fmt"
)

func main() {
	n1, n2 := 6, 9
	fmt.Println(sumar(n1, n2))

	fmt.Println(sumatoria(23, 5, 6, 7, 8, 9, 10))
	fmt.Println(sumatoria(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func sumar(numero1, numero2 int) (int, bool) {
	return numero1 + numero2, true
}

//Asiendo uso de cliclos for-each, sprinf y del operador 3 puntos
func sumatoria(numeros ...int) (int, string) {
	var total int
	for val := range numeros {
		total += val
	}
	return total, fmt.Sprint("El total de la sumatoria es ", total)
}
