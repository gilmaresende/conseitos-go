package main

func main() {

	var soma = 1 + 2
	var subtracao = 1 - 2
	var divisao = 10 / 4
	var miltiplcacao = 10 * 5
	var restoDivisao = 10 % 2

	println("soma", soma)
	println("subtracao", subtracao)
	println("divisao", divisao)
	println("miltiplcacao", miltiplcacao)
	println("restoDivisao", restoDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	var soma2 int16 = numero1 + numero2
	println(soma2)

	var varialve1 string = "String"
	varialve2 := "fd"

	println(varialve1, varialve2)

	println(1 > 2)

	verdadeiro, falso := true, false

	println(verdadeiro && falso)
	println(verdadeiro || falso)

	println(!verdadeiro || falso)

	numero := 10
	println(numero)
	numero++
	println(numero)
	numero += 15
	println(numero)
	numero--
	println(numero)
	numero -= 15
	println(numero)

}