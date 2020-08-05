package main

import (
	"fmt"
)

func main() {
	//Channel send only
	ca := make(chan<- int, 2)

	ca <- 44
	ca <- 66
	//fmt.Println(<-ca)
	//fmt.Println(<-ca)
	//-----------------------------
	fmt.Println("-------")
	fmt.Printf("%T\n", ca)
	//Channel recive only
	ca2 := make(<-chan int, 2)

	ca2 <- 44
	ca2 <- 66
	fmt.Println(<-ca2)
	fmt.Println(<-ca2)
	//-----------------------------
	fmt.Println("-------")
	fmt.Printf("%T\n", ca)
}
