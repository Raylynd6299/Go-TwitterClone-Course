package main

import "fmt"

func main() {
	var numero1, numero2 int
	fmt.Println("Ingrese numero 1:")
	fmt.Scanf("%d", &numero1)
	fmt.Println("Ingrese numero 2:")
	fmt.Scanf("%d", &numero2)

	fmt.Println(numero1 + numero2)
}
