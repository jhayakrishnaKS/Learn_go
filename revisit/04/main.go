package main

import (
	"fmt"
)

// Todo represents a single todo item
type Todo struct {
	Title    string
	Complete bool
}

// AddTodo adds a new todo to the list
func AddTodo() Todo {
	var title string
	fmt.Print("Enter todo title: ")
	fmt.Scanln(&title)
	return Todo{Title: title}
}

// CompleteTodo marks a todo as complete
func (t *Todo) CompleteTodo() {
	t.Complete = true
}

func main() {
	var todo Todo

	fmt.Println("Welcome to Todo App!")

	for {
		fmt.Println("\n1. Add Todo\n2. Complete Todo\n3. Quit")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			todo = AddTodo()
			fmt.Println("Todo added successfully.")
		case 2:
			if todo.Title != "" {
				todo.CompleteTodo()
				fmt.Println("Todo marked as complete.")
			} else {
				fmt.Println("No todo to mark as complete.")
			}
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
