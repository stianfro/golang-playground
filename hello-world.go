package main

import (
	"fmt"
	"math/cmplx"
)

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
}
