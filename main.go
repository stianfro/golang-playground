package main

import "fmt"

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

// Page 33
func main() {
	gohan := &Saiyan{
		Name:  "Gohan",
		Power: 1000,
		Father: &Saiyan{
			Name:   "Goku",
			Power:  9001,
			Father: nil,
		},
	}
	fmt.Println(gohan.Father.Name)
}
