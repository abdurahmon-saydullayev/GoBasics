package main

import (
	"fmt"
	"math"
)

func square() {
	var a,b,c float64
	
	fmt.Printf("a ni kirit:")
	fmt.Scan(&a)
	fmt.Printf("b ni kirit:")
	fmt.Scan(&b)
	fmt.Printf("c ni kirit:")
	fmt.Scan(&c)

	result:=b*b-4*a*c

	defer func(){
		if r:=recover();r!=nil{
			fmt.Println("Recovered in",result)
		}
	}()

	if result<0{
		fmt.Println("Bo'sh to'plam")
		panic(result)
	}

	if result==0{
		fmt.Println("Bitta yechim")
		func1(a,b,c)
	}

	if result>0{
		fmt.Println("Ikta yechim")
		func2(a,b,c)
	}
}

func func1(a,b,c float64){
	x:=-b/2*a
	fmt.Println("x:",x)
}

func func2(a,b,c float64){
	x1:=(-b+(math.Sqrt(b*b-4*a*c)))/(2*a)
	x2:=(-b-(math.Sqrt(b*b-4*a*c)))/(2*a)
	fmt.Printf("x1:%v\nx2:%v",x1,x2)
}