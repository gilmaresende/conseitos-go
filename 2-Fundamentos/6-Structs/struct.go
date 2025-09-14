package main

import "fmt"

type Usuario struct {
	nome     string
	idade    uint8
	endereco Endereco
}

type Endereco struct {
	logradouro string
	numero     string
}

func main() {
	fmt.Println("Arquivo structs")
	end1 := Endereco{"Rua dos bobos", "0"}

	var u Usuario
	u.nome = "Fulano"
	u.idade = 29
	u.endereco = end1
	fmt.Println(u)

	u2 := Usuario{"Ciclano", 29, end1}
	fmt.Println(u2)

	u3 := Usuario{nome: "Beltrano"}
	fmt.Println(u3)

}
