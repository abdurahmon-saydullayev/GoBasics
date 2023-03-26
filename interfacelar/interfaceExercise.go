package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var book []Book

type BookData interface{
	GetBooks()[]Book
	AddBook(Book)
	RemoveBokok(Book)
	Update(Book)
	GetBookById(Book)
	GetByCategory(Book)
}

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	Status   string `json:"status"`
	Price    int    `json:"price"`
	Period   int    `json:"period"`
	Category string `json:"category"`
	Page     int    `json:"page"`
}

type   struct{
	books Book
}

func (b BookData)GetBook()[]Book{
	return b.GetBooks()
}

func(c Book)GetBooks()[]Book{
	data:=ReadFileFROMJSON()
	return data
}

func printData(s BookData){
	fmt.Println("All bokks:",s.GetBooks)
}

func interfaceExercise() {
	d:=printData{BookData:da}
	fmt.Println(d)
}

func ReadFileFROMJSON()[]Book{
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return book
	}

	err = json.Unmarshal(data, &book)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return book
	}
	return book
}