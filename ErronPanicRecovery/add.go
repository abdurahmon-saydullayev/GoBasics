package main

import "fmt"

func add(num1,num2 string)string{

	defer func(){
		if r:=recover();r!=nil{
			fmt.Println("Recovered in",r)
		}
	}()

	if num1==""{
		fmt.Println("Empty")
		panic(fmt.Sprintf("%v",num1))
	}
	if num2==""{
		fmt.Println("Empty")
		panic(fmt.Sprintf("%v",num2))
	}

	return num1 +num2
}

