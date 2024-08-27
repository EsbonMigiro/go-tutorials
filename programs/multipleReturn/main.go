package main

import (
	"fmt"
	"strings"
)

func moreReturns(sentence string) (string, string) {
	sentenceCaps := strings.ToUpper(sentence)
	sentenceSplitted := strings.Split(sentenceCaps, " ")

	var initials []string

	for _, value := range sentenceSplitted {
		initials = append(initials, value[:1])
	}
	if len(initials) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"

}
func main() {
	a1, a2 := moreReturns("Michael Coding")
	fmt.Println("a1", a1, "a2", a2)

	a3, a4 := moreReturns("Michael")
	fmt.Println("a3", a3, "a4", a4)
}
