package main

import (
	"fmt"
)

var result int

func main() {
	fmt.Println("inicio")
	result = operacionesmidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesmidd(restar)(2, 3)
	fmt.Println(result)
	result = operacionesmidd(multiplicar)(2, 3)
	fmt.Println(result)
}
func operacionesmidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}
