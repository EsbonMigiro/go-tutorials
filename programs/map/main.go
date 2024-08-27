package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup":  10.2,
		"salad": 69.3,
		"sossi": 58.3,
	}
	menu["chapo"] = 10.59
	menu["amon"] = 10.548

	fmt.Println(menu)
	fmt.Println(menu["salad"])

	//loop

	for key, value := range menu {
		fmt.Println(key, ":", value)
	}

	phoneBook := map[int]string{
		2: "Ahmed",
		4: "Jonte",
		6: "Jess",
	}
	fmt.Println(phoneBook)
	fmt.Println(phoneBook[6])

	phoneBook[2] = "Michael"
	fmt.Println(phoneBook)

}
