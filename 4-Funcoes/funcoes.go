package main

import "fmt"

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println("Função F:" + txt)
		return "Função F:" + txt
	}

	f("oi")

	resultadosSoma, resultadoSub := calculosMatematicos(1, 3)
	fmt.Println(resultadosSoma)
	fmt.Println(resultadoSub)

	resultadosSoma2, _ := calculosMatematicos(1, 3)
	fmt.Println(resultadosSoma2)
}

func somar(i1, i2 int) int {
	return i1 + i2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}
