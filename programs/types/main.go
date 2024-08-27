package main

import "fmt"

func updateName(nam string) string {
	nam = "Jonte"

	return nam
}
func updateDrinks(drin map[string]float64) {
	drin["coffee"] = 38.7
}

func main() {
	name := "David"

	name = updateName(name)
	fmt.Println(name)

	drinks := map[string]float64{
		"tea":   38.3,
		"cocoa": 58,
	}
	updateDrinks(drinks)
	fmt.Println(drinks)

}
