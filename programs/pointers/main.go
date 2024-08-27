package main

import "fmt"

func updateName(nam *string) {
	*nam = "Jonte"
}
func updateDrinks(drin map[string]float64) {
	drin["coffee"] = 38.7
}

func main() {
	name := "David"

	m := &name

	// fmt.Println("mem location", m)
	// fmt.Println("mem value: ", *m)
	fmt.Println("name", name)
	updateName(m)
	fmt.Println("name", name)

}
