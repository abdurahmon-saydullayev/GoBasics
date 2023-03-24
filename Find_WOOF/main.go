package main

import "fmt"

func main() {
	var str string
	fmt.Printf("Enter string:")
	fmt.Scan(&str)
	counterW:=0
	counterO:=0
	counterF:=0
	for _,v:= range str{
		if v=='w'{
			counterW++
		}
		if v=='o'{
			counterO++
		}
		if v=='f'{
			counterF++
		}
	}

	if counterW<counterF && counterO>=counterW*2 && counterW>0{
		fmt.Printf("%vta woof\n",counterW)
	} else if counterF<counterW &&counterO>=counterF*2 && counterF>0{
		fmt.Printf("%vta woof\n",counterF)
	} else if counterF==counterW &&counterO>=counterF{
		fmt.Printf("%vta woof\n",counterF)
	}else{
		fmt.Println("None")
	}
}