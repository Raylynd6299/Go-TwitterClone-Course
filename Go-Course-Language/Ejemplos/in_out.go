package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var numero1, numero2 int
	var leyenda string
	fmt.Println("Ingrese numero 1:")
	fmt.Scanf("%d", &numero1)

	fmt.Println("Ingrese numero 2:")
	fmt.Scanf("%d", &numero2)

	fmt.Println(numero1 + numero2)

	fmt.Println("Ingrese una Descripcion")

	// Esta es otra forma de ingresar datos, aunque un poco mas complicada
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text()
	}
	resul := numero1 + numero2
	fmt.Println(resul, leyenda)
}
