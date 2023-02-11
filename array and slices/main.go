package main

import "fmt"

func main() {
	// fmt.Println("Hello World!")
	// array

	var age = [3]int{12, 34, 56}
	name := [4]string{"Japan", "Bangladesh", "UAE", "Canada"}
	fmt.Println(age, len(age))
	fmt.Println(name, len(name))

	// slices (use arrays uder the hood)
	scores := []int{123, 675, 577}
	scores[0] = 100
	scores = append(scores, 890)

	fmt.Println(scores, len(scores))

	// slices ranges
	rangeOne := name[1:2]
	rangeTwo := name[3:]
	rangeThree := name[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

}
