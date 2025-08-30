package main

import (
	"errors"
	"fmt"
)

func main() {

	var numero8 int8 = 80

	fmt.Println(numero8)

	var numero16 int16 = 16000

	fmt.Println(numero16)

	var numero32 int32 = 320000000

	fmt.Println(numero32)

	var numero64 int64 = 6400000000000000000

	fmt.Println(numero64)

	var intArquitetura int = 1000000000000000000

	fmt.Println(intArquitetura)

	var decimal32 float32 = 123000000000000000000000000000000000000.444444444444

	fmt.Println(decimal32)

	var decimal64 float64 = 123450000000000000000000000000000000000000000000000000000000000000000000000000000.222222222222222

	fmt.Println(decimal64)

	var texto string = "conteudo da string"

	fmt.Println(texto)

	texto2 := "Texto2"

	fmt.Println(texto2)

	char := 'B'

	fmt.Println(char)

	var variavelTextoDeValorZero string

	fmt.Println(variavelTextoDeValorZero)

	var variavelNumeralDeValorZero int16

	fmt.Println(variavelNumeralDeValorZero)

	var booleanExemplo bool

	fmt.Println(booleanExemplo)

	var errorExemploVazio error

	fmt.Println(errorExemploVazio)

	var errorInscia error = errors.New("Erro interno")

	fmt.Println(errorInscia)

}
