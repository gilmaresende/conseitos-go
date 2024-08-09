package main

import "fmt"

func main() {
	println("Ponteiro")

	var varialve1 int = 10
	var varialve2 int = varialve1

	fmt.Println(varialve1, varialve2)

	varialve1++

	fmt.Println(varialve1, varialve2)

	var varialve3 int = 100
	var ponteiro *int

	ponteiro = &varialve3

	fmt.Println(varialve3, *ponteiro)

*ponteiro++
	
	fmt.Println(varialve3, *ponteiro)

}