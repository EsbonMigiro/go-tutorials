package main

import "fmt"

func main() {
	age := 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	if age < 39 {
		fmt.Println("age is less", age)
	} else if age < 40 {
		fmt.Println("age is less", age)
	} else {
		fmt.Println("age is ", age)

	}

	names := []string{"Michael", "John", "Dan", "David", "Pato", "James", "Elvis"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("One: ", index)
			continue
		}
		if index == 3 {
			fmt.Println("breaking at ", value)
			break
		}
		fmt.Println(index)
	}

}
