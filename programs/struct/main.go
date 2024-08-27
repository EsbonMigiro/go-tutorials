package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	str, err := r.ReadString('\n')
	str = strings.TrimSpace(str)

	return str, err

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	// fmt.Print("create bill: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name)
	name, _ := getInput("create your bill: ", reader)

	b := newBill(name)
	fmt.Println("created ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("chose your options: (a: add bill, s: save bill, t)", reader)
	// fmt.Println(opt)
	switch opt {
	case "a":
		name, _ := getInput("enter the item name: ", reader)
		price, _ := getInput("enter the item price: ", reader)

		val, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("price must be a number")
			promptOptions(b)
		}
		b.addItems(name, val)
		fmt.Println(name, val)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Println("you saved", b.name)
	case "t":
		tip, _ := getInput("enter the item price: ", reader)

		val, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("price must be a number")
			promptOptions(b)
		}
		b.updateTip(val)
		fmt.Println("added tip", val)
		promptOptions(b)
	default:
		fmt.Println("Invalid Option")
		promptOptions(b)
	}
}

func main() {
	// newOneBill := newBill("Me")
	// fmt.Println(newOneBill)
	// newOneBill.updateTip(20)

	// newOneBill.addItems("onion", 25.39)
	// newOneBill.addItems("orange", 235.39)
	// newOneBill.addItems("tomatoes", 205.39)

	// billInfo := newOneBill.format()
	// fmt.Println(billInfo)'

	myBill := createBill()
	promptOptions(myBill)
	// fmt.Println(myBill)

}
