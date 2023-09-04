package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("arrays e slices")

	var array1 [5]string
	array1[0] = "posicao 1"
	fmt.Println(array1)

	array2 := [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)
	fmt.Println("------* SLICE *------")

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 25)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "changed"
	fmt.Println(slice2)

	//arrays internos
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

}
