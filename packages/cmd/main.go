package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gitlab.com/i.abdukhoshimov/golang_bootcamp/new-project/storage/repo"
)

func main() {
	// Open the JSON file
	file, err := os.Open("students.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the JSON data into a slice of Student structs
	var students []repo.Student
	err = json.NewDecoder(file).Decode(&students)
	if err != nil {
		log.Fatal(err)
	}

	// Create a School struct to hold the student data
	school := &School{students}

	// Example usage
	fmt.Println(school.GetStudents())

	student, err := school.GetStudentByID(2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(student)

	err = school.AddStudent(Student{
		ID:      4,
		Name:    "John Doe",
		Grade:   9,
		Section: "A",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(school.GetStudents())

	err = school.RemoveStudentByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(school.GetStudents())
}
