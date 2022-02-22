package main

type Saiyan struct {
	Name   string
	Power  int
	Father *Saiyan
}

func main() {
	goku := &Saiyan{Name: "Goku", Power: 9001}

	extractPowers(goku)
}

func extractPowers(saiyans []*Saiyan) []int {
	powers := make([]int, len(saiyans))
	for index, saiyan := range saiyans {
		powers[index] = saiyan.Power
	}
	return powers
}
