package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero uint8
}

func main() {
	println("Arquivo struct")

	var usuario1 usuario
	usuario1.nome = "Gil"
	usuario1.idade = 28
	fmt.Println(usuario1)

	endereco1 := endereco{"Rua dos bobos", 0}

	usuario2 := usuario{"Gil2", 28, endereco1}
	fmt.Println(usuario2)

	usuario3 := usuario{nome:"Ana"}
	fmt.Println(usuario3)

}