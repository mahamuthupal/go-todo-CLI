package main

import (
    "flag"
    "fmt"
    "os"

	"goCLI/todo"
)

const todoFileName = ".todo.json"

func main() {

	// parsing command line flags, these are pointers
	task := flag.String("task", "", "Task to be included in the todolist")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()
	// pointer to todo.List by address operator to extract the address of an empty instance of todo.List
	l := &todo.List{}

	//calling Get method from todo.go file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Decide what to do based on arguments provided
	switch {
	case *list:
		// list current to do items
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	// to verify the if complete flag is set with value more than 0 (default)
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	// verify if task flag is set with different than empty string
	case *task != "":
		l.Add(*task)
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// print an error msg
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}

}
