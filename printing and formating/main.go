package main

import "fmt"

func main() {
	age := 26
	name := "Jahidul Hasan Zahid"
	// print
	// \n is used for a new line
	fmt.Print("Hello, ")
	fmt.Print("World \n")
	fmt.Print("New Line started \n")
	fmt.Print("Hello Bangladesh! \n")
	fmt.Println("Hello World!")
	fmt.Println("Hello Dhaka")

	fmt.Println("Hello, My name is ", name, "and my age is ", age)

	// printf formating string
	fmt.Printf("Hello, My age is %v and my name is %v \n", age, name)
	fmt.Printf("Yy name is %q and my age is %q \n", name, age)
	fmt.Printf("Age is of type %T \n", age)
	fmt.Printf("Your score is %f points \n", 23.65687)
	fmt.Printf("Your score is %0.2f point \n", 23.564645)

	// Sprintf save formatted strings
	var str = fmt.Sprintf("My name is %v and my age is %v \n", name, age)
	fmt.Println("The Saved string is: ", str)
}
