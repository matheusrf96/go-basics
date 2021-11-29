package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr1 [5]string
	arr1[0] = "Posição 1"
	fmt.Println(arr1)

	arr2 := [5]string{"index1", "index2", "index3", "index4", "index5"}
	fmt.Println(arr2)

	arr3 := [...]int{1, 2, 3}
	fmt.Println(arr3)

	slice := []int{11, 12, 13, 14, 15}
	fmt.Println(slice)

	slice = append(slice, 16)
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(arr3))

	slice2 := arr2[1:3]
	fmt.Println(slice2)

	arr2[1] = "altered index"
	fmt.Println(slice2)

	fmt.Println("-----")

	// Internal arrays

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // length
	fmt.Println(cap(slice3)) // capacity
	fmt.Println()

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	fmt.Println()

	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
	fmt.Println()

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	fmt.Println()

	slice4 = append(slice4, 6)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	fmt.Println()
}
