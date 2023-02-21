package main

import "fmt"

func main() {
	// math := 1
	// for math < 10 {
	// 	fmt.Println("Math Number is: ", math)
	// 	math++
	// }

	// for bangla := 1; bangla < 10; bangla++ {
	// 	fmt.Println("Bangla Number is: ", bangla)
	// }

	names := []string{"Zahid", "Rashed", "Rafan", "Rubel", "Sagor"}
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for index, value := range names {
	// 	fmt.Printf("The value at index %v is %v \n", index, value)
	// }

	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
	}

	fmt.Println(names)

}
