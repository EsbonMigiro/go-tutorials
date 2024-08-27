package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name: name,
		// items: map[string]float64{"pie": 58.00, "salad": 35.83},
		items: map[string]float64{},

		tip: 0,
	}
	return b

}

func (b *bill) format() string {
	fs := fmt.Sprintf("The Bill: \n")
	var total float64 = 0

	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", key+":", value)
		total += value
	}
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total", b.tip)

	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total", total)

	return fs

}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItems(name string, price float64) {
	b.items[name] = price
}

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("saved successfully")
}
