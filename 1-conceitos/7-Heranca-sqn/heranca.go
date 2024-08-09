package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int8
	autura    uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	println("Herança")
	pessoa1 := pessoa{"Jão", "Pedro", 20, 178}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "CC", "UIT"}
	fmt.Println(estudante1)

	fmt.Println("Altura do", estudante1.nome, "é", estudante1.autura)


}