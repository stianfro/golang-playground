package main

import "fmt"

type Saiyan struct {
	*Person
	Power int
}

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

func (s *Saiyan) Introduce() {
	fmt.Printf("Hi, I'm %s. Ya!\n", s.Name)
}

// Page 33
func main() {
	goku := &Saiyan{
		Person: &Person{"Goku"},
		Power:  9001,
	}
	goku.Introduce()
	fmt.Println(goku.Name)
	fmt.Println(goku.Person.Name)
}
