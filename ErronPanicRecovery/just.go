package main

import "fmt"

func sillySusan() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Susan recovered from the panic")
			fmt.Println("silly susan handles the panic gracefully")
		}
	}()
	fmt.Println("silly susan called")
	panickingPeter()
	fmt.Println("silly susan finished")
}

func panickingPeter() {
	fmt.Println("panicking peter called")
	panic("oh no")
}