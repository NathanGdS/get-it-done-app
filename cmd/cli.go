package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	add := flag.Bool("add", false, "add a new todo")
	complete := flag.Int("complete", 0, "mark a todo as completed")
	del := flag.Int("del", 0, "delete a todo")
	list := flag.Bool("list", false, "list all todos")

	flag.Parse()

	switch {
	case *add:
		fmt.Println("add todo")
	case *complete > 0:
		fmt.Println("complete a given todo")
	case *del > 0:
		fmt.Println("delete a given todo")
	case *list:
		fmt.Println("list all todos")
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}

}
