package main

import "fmt"

func main() {
	// var num1, num2 string
	// fmt.Printf("Number 1:")
	// fmt.Scan(&num1)
	// fmt.Printf("Number 1:")
	// fmt.Scan(&num2)

	fmt.Println("cascading panics")
	sillySusan()
	fmt.Println("again")

	add("45","56")
	fmt.Println("recovered")
	add("","34")
}