package main

import "fmt"

func main() {
	age := 24
	name := "Michael"
	//print
	fmt.Print("Hello Michael")
	fmt.Print("Hello Michael \n")
	fmt.Print("Hello Michael \n")
	// println

	fmt.Println("my name is ", name, "I am ", age, "old")

	//printf(formated)

	fmt.Printf("my name is %v I am %v old \n", name, age)
	fmt.Printf("my name is %q I am %q old \n", name, age)
	fmt.Printf("age is of tyepe %T \n", age)

	fmt.Printf("The value is %f \n", 255.55)
	fmt.Printf("The value is %0.1f \n", 225.55)

	var str = fmt.Sprintf("my name is %v I am %v old \n", name, age)
	fmt.Println("string is : ", str)
}
