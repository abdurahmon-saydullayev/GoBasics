package main

import "fmt"

func main() {
	var a,b int
	a=3
	b=2
	pointer(&a,&b)
}
func pointer(a,b *int){
	c:=*a
	*a=*b
	*b=c	
	fmt.Println(*a,*b)
}