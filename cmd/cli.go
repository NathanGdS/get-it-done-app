package main

import (
	"flag"
	"fmt"

	"github.com/nathangds/get-it-done-app/constants"
	"github.com/nathangds/get-it-done-app/handlers"
)

func main() {
	add := flag.String("add", "", "add a todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	del := flag.Int("del", 0, "delete a todo")
	list := flag.Bool("list", false, "list all todos")
	target := flag.String("target", constants.DEFAULT_FILENAME, "target filename to store todos")

	flag.Parse()

	switch {
	case *add != "":
		handlers.AddTodo(*add, *target)
	case *complete > 0:
		handlers.CompleteTodoById(*complete, *target)
	case *del > 0:
		handlers.DeleteTodoById(*del, *target)
	case *list:
		handlers.ListTodosCli(*target)
	default:
		fmt.Println("Invalid option, loading all todos...")
		handlers.ListTodosCli(*target)
	}

}
