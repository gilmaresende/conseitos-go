package main

import "fmt"

func main() {
	aritimeticos()
	atribuicao()
	relacionais()
	logicos()
	unarios()
	tenario()
}

func tenario() {
	//não existe
	//texto :=1>5? "V":"F"
	//tem que usar o velho if
	var text string
	if 1 > 5 {
		text = "V"
	} else {
		text = "F"
	}
	print(text)
}

func unarios() {
	n1 := 10
	println(n1)
	n1++
	println(n1)
	n1--
	println(n1)
	n1 += 5
	println(n1)
	n1 -= 3
	println(n1)

	n1 /= 3
	println(n1)

	n1 *= 4
	println(n1)

}

func logicos() {
	verdadeiro, falso := false, true
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
}

func relacionais() {
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 != 2)
}

func atribuicao() {
	var vari1 string = "String"
	vari2 := "String2"
	fmt.Println(vari1, vari2)
}

func aritimeticos() {
	soma := 1 + 1
	substracao := 1 - 1
	divisao := 2 / 2
	miltiplicacao := 2 * 2
	restoDivisao := 10 % 3

	fmt.Println(soma, substracao, divisao, miltiplicacao, restoDivisao)

	var numero1 int16 = 10

	//se for int32 não deixa somar
	var numero2 int16 = 25

	soma2 := numero1 + numero2
	fmt.Println(soma2)
}
