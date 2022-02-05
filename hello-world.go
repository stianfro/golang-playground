package main

import (
	"encoding/json"
	"fmt"
	"math/cmplx"
	"time"
)

type response struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

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

func f(from string) {
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println(from, ":", i)
	}
}

func fizzBuzz(i int) {
	if i%3 == 0 && i%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if i%3 == 0 {
		fmt.Println("Fizz")
	} else if i%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(i)
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

	// For loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// While loop
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)

	// Infinite loop
	// for {
	// 	do something
	// }

	// Each loop
	kvs := map[string]string{"name": "Stian", "lastname": "Fro"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// Each with index
	arr := []string{"a", "b", "c", "e"}
	for index, value := range arr {
		fmt.Println(index, value)
	}

	// Convert to json
	mapA := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapA)
	fmt.Println(string(mapB))

	// Convert from json
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res.Fruits)

	// Concurrency
	go f("goroutine")

	f("main")
	fmt.Println("done")

	// Fizzbuzz
	count := 20
	for i := 10; i < count; i++ {
		fizzBuzz(i)
	}
}
