package main

import (
	"fmt"
	"math"
)

func sayGreetings(name string) {
	fmt.Printf("Goodmorning %v \n", name)
}

func additionBy(name string, values []int) {
	sum := values[0] + values[1]

	fmt.Printf("summation as per %v is %v \n", name, sum)
}
func greetingPeople(names []string, val []int, functioner func(string, []int)) {
	for _, value := range names {
		functioner(value, val)
	}
}
func areaOfCircle(r float64) float64 {
	return math.Pi * r * r
}
func main() {
	// sayGreetings("Mario")
	// sayGreetings("Omosa")
	// additionBy("Michael", []int{20, 50})
	// greetingPeople([]string{"Davido", "Jonte"}, []int{20, 50}, additionBy)

	a1 := areaOfCircle(2)
	a2 := areaOfCircle(7)

	fmt.Printf("areaOne: %0.3f, areaTwo %0.1f \n", a1, a2)

}
