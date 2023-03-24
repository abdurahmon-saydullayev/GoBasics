package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	
	var people [] Person
	err = json.Unmarshal(data, &people)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	for _, person := range people {
		fmt.Println(person.Name, person.Age)
	}

	people = append(people, Person{Name: "Sahob", Age: 23}, Person{Name: "Samandar", Age: 30})

	
	for i, person := range people {
		if person.Name == "Diyorbek" {
			people[i].Age = 40
		}
	}

	for i,person := range people{
		if person.Name=="John"{
			people=append(people[:i],people[i+1:]...)

		}
	}
    
	data, err = json.Marshal(people)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	err = ioutil.WriteFile("data.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}

}
