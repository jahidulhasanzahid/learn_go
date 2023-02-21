package main

import (
	"fmt"
	"strings"
)

func getInit(name string) (string, string) {
	s := strings.ToUpper(name)
	names := strings.Split(s, " ")

	var init []string
	for _, value := range names {
		init = append(init, value[:1])
	}

	if len(init) > 1 {
		return init[0], init[1]
	}

	return init[0], "_"
}
func main() {
	fn1, sn1 := getInit("jahidul hasan")
	fmt.Println(fn1, sn1)

	fn2, sn2 := getInit("Sagor")
	fmt.Println(fn2, sn2)
}
