package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type BookData interface {
	GetBooks() []Book
	AddBook(Book)
	RemoveBook(Book)
	UpdateBook(Book)
	GetBookById(Book)
	GetBookByCategory(Book)
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
type BookLibrary struct {
	Books []Book
}

func (b *BookLibrary) GetBooks() []Book {
	return b.Books
}

func (b *BookLibrary) AddBook(book Book) []Book {
	b.Books = append(b.Books, book)
	return b.Books
}

func (b *BookLibrary) GetBookById(book Book)[]Book{
	for _,v:= range b.Books{
		if v.ID==book.ID{
			fmt.Println(v)
		}
	}
	return b.Books
}

func (b *BookLibrary) RemoveBook(book Book)[]Book{
	for i,v:= range b.Books{
		if v.ID==book.ID{
			b.Books= append(b.Books[:i],b.Books[i+1:]...)
			break
		}
	}
	return b.Books
}

func (b *BookLibrary) UpdateBook(book Book)[]Book{
	for _,v:= range b.Books{
		if v.ID==book.ID{
			v.Title="Abdurahmon"
			break
		}
	}
	return b.Books
}

func (b *BookLibrary) GetBookByCategory(book Book){
	for _,v:= range b.Books{
		if v.Category==book.Category{
			fmt.Println(v)
		}
	}
}

func main() {
	data,_:=ioutil.ReadFile("books.json")
	var book []Book
	_=json.Unmarshal(data,&book)

	library := BookLibrary{book}

	
	// fmt.Println(library.GetBooks())

	// fmt.Println(library.AddBook(Book{
	//     ID :4,    
    //     Title: "me",
    //     Author  :"me",
    //     Year:2020,
    //     Status :"done",
    //     Price :150,  
    //     Period  :12,
    //     Category:"Fiction",
    //     Page :150,
	// }))

	// fmt.Println(library.RemoveBook(Book{
	// 	ID :3,    
    //     Title: "me",
    //     Author  :"me",
    //     Year:2020,
    //     Status :"done",
    //     Price :150,  
    //     Period  :12,
    //     Category:"Fiction",
    //     Page :150,
	// }))

	// fmt.Println(library.UpdateBook(Book{
	// 	ID :2,    
    //     Title: "me",
    //     Author  :"me",
    //     Year:2020,
    //     Status :"done",
    //     Price :150,  
    //     Period  :12,
    //     Category:"Fiction",
    //     Page :150,
	// }))

	library.GetBookById(Book{
		ID: 3,
	})

	// library.GetBookByCategory(Book{
	// 	Category: "Fiction",
	// })
}