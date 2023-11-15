package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nathangds/get-it-done-app/handlers"
)

func main() {
	add := flag.String("add", "", "add a todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	del := flag.Int("del", 0, "delete a todo")
	list := flag.Bool("list", false, "list all todos")

	flag.Parse()

	switch {
	case *add != "":
		handlers.AddTodo(*add)
	case *complete > 0:
		fmt.Println("complete a given todo")
	case *del > 0:
		fmt.Println("delete a given todo")
	case *list:
		handlers.ListTodosCli()
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}

}
