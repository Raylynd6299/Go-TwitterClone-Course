package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	leoArchivo()
	leoArchivo2()
	escribirArchivo()
	escribirArchivo2()
}

func leoArchivo() {
	archivo, err := ioutil.ReadFile("./prueba.txt")
	if err != nil {
		log.Fatal("Error al obtener el archivo")
	}
	fmt.Println(string(archivo))
}

func leoArchivo2() {
	archivo, err := os.Open("./prueba.txt")
	defer archivo.Close()
	if err != nil {
		log.Fatal("Error al obtener el archivo")
	}

	scanner := bufio.NewScanner(archivo)

	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("Linea > ", registro)
	}
}

func escribirArchivo() {
	archivo, err := os.Create("./prueba2.txt")
	defer archivo.Close()
	if err != nil {
		log.Fatal("Error al obtener el archivo")
	}
	fmt.Fprintln(archivo, "Logramos exscribir una linea en un archivo nuevo")
}

func escribirArchivo2() {
	filename := "./prueba2.txt"
	if Append(filename, "Esta es una segunda linea") == false {
		log.Fatal("Error al escribir en el archivo")
	}
}

func Append(archivo string, texto string) bool {
	arch, err := os.OpenFile(archivo, os.O_WRONLY|os.O_APPEND, 0744)
	defer arch.Close()
	if err != nil {
		log.Fatal("Error al obtener el archivo")
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		log.Fatal("Error escribir en el archivo")
	}
	return true
}
