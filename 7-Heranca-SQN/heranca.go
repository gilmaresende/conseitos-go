package main

import "fmt"

type Pessoa struct {
	nome      string
	documento string
	idade     int8
	altura    uint8
}

type Estudante struct {
	Pessoa
	curso     string
	faculdade string
}

func main() {
	println("Heranca")

	p1 := Pessoa{"Fulano", "125", 50, 180}
	fmt.Println(p1)

	var p2 Estudante = Estudante{p1, "Computacao", "UIT"}
	fmt.Println(p2)
	fmt.Println(p2.nome)
}
