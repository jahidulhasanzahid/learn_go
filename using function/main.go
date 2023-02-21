package main

import (
	"fmt"
	"math"
)

// func sayGoodmorning(name string) {
// 	fmt.Printf("Good Morning %v \n", name)
// }
// func sayGoodnight(name string) {
// 	fmt.Printf("Good Night %v \n", name)
// }
// func cycleNames(names []string, function func(string)) {
// 	for _, value := range names {
// 		function(value)
// 	}
// }

func cycleArea(number float64) float64 {
	return math.Pi * number * number
}
func main() {
	// sayGoodmorning("Jahidul")
	// sayGoodmorning("Rashed")
	// sayGoodnight("Zahid Hasan")
	// sayGoodnight("Sagor Banik")
	// cycleNames([]string{"Jahidul", "Rashed"}, sayGoodmorning)
	// cycleNames([]string{"Jahidul", "Rashed"}, sayGoodnight)
	area1 := cycleArea(10.5)
	area2 := cycleArea(65.98)
	area3 := cycleArea(45.3)

	// fmt.Println(area1, area2, area3)
	fmt.Printf("Circle 1 is %0.3f and circle 2 is %0.3f and circle 3 is %0.3f", area1, area2, area3)
}
