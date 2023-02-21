package main

import "fmt"

func main() {
	// age := 25

	// fmt.Println(age <= 50)
	// fmt.Println(age <= 45)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 55)
	// fmt.Println(age != 50)

	// if age < 45 {
	// 	fmt.Println("Age is less than 45")
	// } else if age < 55 {
	// 	fmt.Println("Age is less than 55")
	// } else {
	// 	fmt.Println("Age is not less than 45")
	// }

	names := []string{"Zahid", "Rashed", "Foysal", "Shakil", "Sagor"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("continuning at position \n", index)
			continue
		}
		if index > 2 {
			fmt.Println("Breaking at position \n", index)
			break
		}
		fmt.Printf("The value at position %v is %v \n", index, value)
	}
}
