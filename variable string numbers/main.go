package main

import "fmt"

func main() {
	// string
	var nameOne string = "Jahidul"
	var nameTwo = "Hasan"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Zahidul"
	nameThree = "jahid"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameour := "Our Name"
	fmt.Println(nameour)

	// ints
	var ageOne int = 26
	var ageTwo = 36
	ageThree := 87

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 27
	var numFour uint16 = 7568
	var numFive uint64 = 9778986654657765648

	fmt.Println(numOne, numTwo, numThree, numFour, numFive)

	// float
	var scoreOne float32 = 45.09
	var scoreTwo float64 = 56575.67

	fmt.Println(scoreOne, scoreTwo)
}
