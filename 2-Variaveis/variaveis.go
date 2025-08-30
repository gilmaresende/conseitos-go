package main

import "fmt"

func main() {
	var variavel1 string = "variavel 1"
	variavel2 := "variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 string = "lelele"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, vavariavel6 := "val5", "val6"

	fmt.Println(variavel5, vavariavel6)

	const const1 string = "constnte 1"

	fmt.Println(const1)

	variavel5, vavariavel6 = vavariavel6, variavel5

	fmt.Println(variavel5, vavariavel6)
}
