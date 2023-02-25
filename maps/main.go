package main

import "fmt"

func main() {
	menu := map[string]float32{
		"soup":           4.65,
		"pie":            7.98,
		"salad":          5.89,
		"toffee puddind": 3.76,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	phonebook := map[int]string{
		4576867545: "Mario",
		7896546455: "Luigi",
		6893453464: "peach",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[4576867545])

	phonebook[4576867545] = "Rashed"
	fmt.Println(phonebook)
}
