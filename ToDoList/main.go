package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type ToDoList interface{
	GetAllTasks()[]Task
	AddTask(Task)
	RemoveTask(Task)
	UpdateTask(Task)
	MarkTask(Task)
}

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
	DueDate     string `json:"dueDate"`
	Priority    string `json:"priority"`
	StartDate   string `json:"startDate"`
}

type TaskList struct {
	Tasks []Task
}

func (b *TaskList) GetAllTasks()[] Task{
	return b.Tasks
}

func (b *TaskList) AddTask(task Task){
	b.Tasks=append(b.Tasks, task)
	fmt.Println(b.Tasks)
}

func (b *TaskList) RemoveTask(task Task){
	for i,v := range b.Tasks{
		if v.Title==task.Title{
			b.Tasks=append(b.Tasks[:i], b.Tasks[i+1:]...)
		}
	}
	fmt.Println(b.Tasks)
}

func (b *TaskList) MarkTask(task Task){
	for _,v:= range b.Tasks{
		if v.Title==task.Title{
			task.Completed=true
		}
	}
	fmt.Println(b.Tasks)
}

func (b *TaskList) UpdateTask(task Task){
	for _,v := range b.Tasks{
		if v.Title==task.Title{
			task.Completed=true
			task.Description="done"
			task.DueDate="12.03"
			task.ID=len(b.Tasks)+1
			task.Priority="high"
			task.StartDate="11.03"
			task.Title="changed"
		}
	}
	fmt.Println(b.Tasks)
}

func main() {
	data,_:=ioutil.ReadFile("task.json")
	var task []Task

	_=json.Unmarshal(data,&task)
	all:=TaskList{task}

	// one:=Task{
	// 	ID          :len(task)+1,
    //     Title       :"Learn Golang",
    //     Completed   : false,
    //     Description :"Deep knowledge for go",
    //     DueDate     :"30.03",
    //     Priority    :"high",
    //     StartDate   :"27.03",
	// }

	// fmt.Println(all.GetAllTasks())
	// all.AddTask(one)
	// all.RemoveTask(Task{
	// 	Title: "Learn Golang",
	// })
	// all.UpdateTask(Task{
	// 	Title: "Learn English",
	// })
	all.MarkTask(Task{
		Title: "Learn English",
	})
}