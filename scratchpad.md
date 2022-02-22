# Scratchpad

This is where I will put old stuff for reference.

## Basic functionality

```go
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

```

## Chapter 2

## Chapter 3

```go
// Initializing slices:
names := []string{"leto", "jessica", "paul"}
checks := make([]bool, 10)
var names []string
scores := make([]int, 0, 20)
```

# Extra

Random knowledge

## Struct slice of structs

This can be useful when making json payloads.

```go
package main

import (
	"encoding/json"
	"fmt"
)

type BagItem struct {
	Name    string    `json:"server"`
	Content []Content `json:"exclusions"`
}
type Content struct {
	GUID        string `json:"guid"`
	Path        string `json:"path"`
	Extension   string `json:"extension"`
	ProcessPath string `json:"processPath"`
}

func main() {
	serverA := &Content{
		GUID:        "357743e9-8fee-4479-997d-66df04f741d5",
		Path:        "bar",
		Extension:   "baz",
		ProcessPath: "babu",
	}
	serverB := &Content{
		GUID:        "22c2e7ee-a792-418e-b588-9271d24b4aff",
		Path:        "bar",
		Extension:   "baz",
		ProcessPath: "babu",
	}
	item := &BagItem{
		Name:    "item-a",
		Content: []Content{*serverA, *serverB},
	}

	payload, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(payload))
}
```

Will produce an output like this:

```json
{
  "server": "server-a",
  "exclusions": [{
    "guid": "357743e9-8fee-4479-997d-66df04f741d5",
    "path": "bar",
    "extension": "baz",
    "processPath": "babu"
  }, {
    "guid": "22c2e7ee-a792-418e-b588-9271d24b4aff",
    "path": "bar",
    "extension": "baz",
    "processPath": "babu"
  }]
}
```
