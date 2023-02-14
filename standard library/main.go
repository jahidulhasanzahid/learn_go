package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var greeting = "Hello there friends!"

	// strings package
	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	// the original value does not change
	fmt.Println("Original Value", greeting)

	// sort package
	ages := []int{98, 67, 12, 34, 45, 67, 25, 78, 45}
	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 45)
	fmt.Println(index)

	names := []string{"D", "A", "C", "B"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.SearchStrings(names, "B"))
}
