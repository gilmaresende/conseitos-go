package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println("Arrays e slices")

	var array1 [5]int

	array1[0] = 10
	fmt.Println(array1)

	array2 := [5]string{"p1", "p2", "p3", "p4"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 5}

	fmt.Println(array3)

	slice := []int{1, 23}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 5)

	fmt.Println(slice)
	fmt.Println(slice[1])

	slice2 := array2[1:3]
	fmt.Println(slice2)

	slice2[1] = "ajustado"

	fmt.Println(array2)

	slice3 := make([]float32, 10, 11)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	fmt.Println(slice3)

	slice3 = append(slice3, 6)
	fmt.Println(slice3)

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
