package main

import (
	"strings"
	"day4/todo/dao/memory"
	"day4/todo/model"
	"fmt"
	"os"
)

var ir = NewInputReader()

var todos = memory.New()

type menuCommand func() error

type menu struct {
	Title   string
	Command menuCommand
}

func newToDoCommand() (err error) {
	fmt.Println("Adding new todo")
	fmt.Print("Enter task details: ")

	task, err := ir.NextString()
	if err != nil {
		return
	}

	todo := model.ToDo{Task: task}

	todos.Add(todo)
	return listToDoCommand()
}

func updateToDoStatusCommand() (err error) {
	fmt.Println("Updating ToDo status")
	fmt.Print("Enter todo ID: ")

	ID, err := ir.NextInt()

	if err != nil {
		return
	}

	todo, err := todos.FindByID(ID)
	if err != nil {
		return
	}

	if (todo == model.ToDo{}) {
		fmt.Printf("no todo by id %v", ID)
		return
	}

	if todo.Completed {
		fmt.Printf("Task '%v' is already completed. Do you want to change it to incomplete? Y/N: ", todo.Task)
		selected := ""
		selected, err = ir.NextString()
		if err != nil {
			return
		}

		if strings.ToUpper(selected) == "Y" {
			todo.Completed = false
			err = todos.UpdateById(todo.ID, todo)
			if err != nil {
				return
			}
		}
	} else {
		fmt.Printf("Task '%v' is incomplete. Do you want to complete it? Y/N: ", todo.Task)
		selected := ""
		selected, err = ir.NextString()
		if err != nil {
			return
		}

		if strings.ToUpper(selected) == "Y" {
			todo.Completed = true
			err = todos.UpdateById(todo.ID, todo)
			if err != nil {
				return
			}
		}
	}
	return listToDoCommand()
}

func deleteToDoCommand() (err error) {
	fmt.Println("Deleting ToDo")
	fmt.Print("Enter todo ID: ")
	ID, err := ir.NextInt()
	if err != nil {
		return
	}
	err = todos.DeleteById(ID)
	return listToDoCommand()
}

func listToDoCommand() (err error) {
	fmt.Println("Listing todos")
	list, err := todos.FindAll()
	if err != nil {
		return
	}

	for _, todo := range list {
		symbol := " "
		if todo.Completed {
			symbol = "\u2713"
		}

		fmt.Printf("%v %v) %v\n", symbol, todo.ID, todo.Task)
	}
	return
}

func exitCommand() (err error) {
	fmt.Println("Exiting...")
	os.Exit(0)
	return
}

func main() {
	defer func() {
		r := recover()

		if r != nil {
			fmt.Println(r)
			os.Exit(1)
		}
	}()

	menus := []menu{
		menu{"Exit", exitCommand},
		menu{"List", listToDoCommand},
		menu{"New", newToDoCommand},
		menu{"Update Status", updateToDoStatusCommand},
		menu{"Delete", deleteToDoCommand},
	}

	for {
		fmt.Println("Menu")
		for i, menu := range menus {
			fmt.Printf("%v) %v\n", i, menu.Title)
		}
		fmt.Print("Choose: ")
		selected, err := ir.NextInt()

		if err != nil {
			panic(err)
		}

		menu := menus[selected]
		err = menu.Command()

		if err != nil {
			panic(err)
		}
	}

}
