package main

import "fmt"

func main() {
	fmt.Println("Arrays e slices")

	var array1[5] string

	array1[0]= "possicao 1"

	fmt.Println(array1)

	array2:=[5]string{"dev1","dev2","dev3","dev4","dev5"}
	fmt.Println(array2)

	//array2[5]= "catatal"

	fmt.Println(array2)

	array3 := [...]int{1,2,3,4,5}
	fmt.Println(array3)


	slice := []int{10,11,12,13,15,16,17}

	fmt.Println(slice)

	slice = append(slice, 18)

	fmt.Println(slice)

	slice2:=array2[1:3]
	fmt.Println(slice2)

	slice2[1] = "POssicao alterada"
	fmt.Println(slice2)




}