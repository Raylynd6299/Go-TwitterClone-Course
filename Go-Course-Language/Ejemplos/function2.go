package main

import "fmt"

func main() {
	//Funcion Anonima
	var Calculo func(int, int) int = func(num1, num2 int) int {
		return num1 + num2
	}

	fmt.Println(Calculo(5, 6))

	//Restar
	Calculo = func(num1, num2 int) int {
		return num1 - num2
	}
	fmt.Println(Calculo(7, 6))

	Operaciones()

	//Closures
	tabladel := 2
	mitabla := Tabla(tabladel)
	for i := 1; i <= 10; i++ {
		fmt.Println(mitabla())
	}

}

func Operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}
	fmt.Println(resultado())
}

//Genera Closures

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}
