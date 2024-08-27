package main

import "fmt"

func main() {
	//
	var ages [3]int = [3]int{10, 20, 30}
	fmt.Println(ages, len(ages))

	var agess = [3]int{10, 20, 30}
	fmt.Println(agess, len(agess))

	names := [4]string{"Obi", "Michael", "alvin", "David"}

	fmt.Println(names, len(names))

	//slice
	var scores = []int{10, 30, 50}
	scores[2] = 60
	scores = append(scores, 80)

	fmt.Println(scores, len(scores))

	//ranges

	rangeOne := scores[0:2]
	rangeTwo := scores[2:4]

	rangeThree := append(rangeTwo, rangeOne...)

	fmt.Println(rangeThree)

}
