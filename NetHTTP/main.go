package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Body struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	res, err:= http.Get("https://jsonplaceholder.typicode.com/posts")
	if err!=nil{
		panic(err)
	}
	defer res.Body.Close()
	
	body,err:=io.ReadAll(res.Body)
	if err!=nil{
		fmt.Println("error while reading res body",err)
	}

	var datas [] Body

	err =json.Unmarshal(body,&datas)
	for _,val:= range datas{
		fmt.Println("ID:",val.ID)
		fmt.Println("UserID:",val.UserID)
		fmt.Println("Title:",val.Title)
		fmt.Println("Body:",val.Body)
		fmt.Println("----------------------------")
	}
}