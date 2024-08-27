package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greetings := "hello there friends"

	fmt.Println(strings.Contains(greetings, "hello"))
	fmt.Println(strings.ReplaceAll(greetings, "hello", "morning"))
	fmt.Println(strings.ToUpper(greetings))
	fmt.Println(strings.Index(greetings, "hello"))
	fmt.Println(strings.Split(greetings, " "))

	//sort

	ages := []int{7, 3, 53, 35, 25, 5, 39, 5, 22}

	sort.Ints(ages)

	fmt.Println("sorted ages: ", ages)

	names := []string{"Michael", "Alvin", "Sam", "David", "Shem"}
	sort.Strings(names)

	fmt.Println(" sorted names", names)

	fmt.Println("search int:", sort.SearchInts(ages, 35))
	fmt.Println("search strings:", sort.SearchStrings(names, "Sam"))

}
