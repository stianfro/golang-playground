package main

import (
	"fmt"
	"math/cmplx"
)

func printNumberString(i int) {
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("none")
	}
}

func main() {
	a := 1
	b := "hello"

	var c bool = true
	var d string = "Commenting in"
	var e float32 = 1.222
	var x complex128 = cmplx.Sqrt(-5 + 12i)

	fmt.Println("[DEBUG]", d, a, b, c, e, x)

	// hash
	hash := map[string]string{"name": "Stian", "lastname": "Froystein"}
	fmt.Println(hash["name"])

	// check if key is present in map
	name := map[string]string{"name": "Stian", "lastName": "Froystein"}
	lastname, ok := name["lastName"]

	if ok {
		fmt.Printf("Last Name: %v \n", lastname)
	} else {
		fmt.Println("Last Name is missing")
	}

	// array
	array := [3]int{1, 2, 3}
	names := [2]string{"Stian", "Aoba"}

	fmt.Println(array, names)

	// slices
	// var f []int
	numbers := []int{1, 2, 3, 4}
	numbers = append(numbers, 2, 4)
	fmt.Println(numbers)

	slice1 := numbers[2:]
	fmt.Println(slice1)

	slice2 := numbers[:3]
	fmt.Println(slice2)

	slice3 := numbers[1:4]
	fmt.Println(slice3)

	array1 := []int{1, 2, 3, 4}
	array2 := make([]int, 4)

	copy(array2, array1)

	// if..else
	num := 9
	if num < 0 {
		fmt.Printf("%d is negative \n", num)
	} else if num < 10 {
		fmt.Printf("%d has 1 digit \n", num)
	} else {
		fmt.Printf("%d has multiple digits \n", num)
	}

	// switch case
	printNumberString(1)
	printNumberString(2)
	printNumberString(3)
	printNumberString(4)
}
