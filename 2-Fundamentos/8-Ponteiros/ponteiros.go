package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel int = 10
	var variavel2 = variavel

	fmt.Println(variavel, variavel2)

	variavel++
	fmt.Println(variavel, variavel2)

	var vari3 int = 100
	var ponteiro *int

	fmt.Println(vari3, ponteiro)

	ponteiro = &vari3

	fmt.Println(vari3, ponteiro)
	fmt.Println(vari3, *ponteiro)

	vari3 = 150

	fmt.Println(vari3, ponteiro)
	fmt.Println(vari3, *ponteiro)

}
