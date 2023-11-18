package handlers

import (
	"fmt"
	"strconv"

	"github.com/alexeyco/simpletable"
	"github.com/nathangds/get-it-done-app/models"
)

func ListTodosCli() {
	var todos = &models.Todos

	if len(*todos) == 0 {
		todoReference := &models.Todo{}
		todoReference.Load("todos.json")
	}

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Description"},
			{Align: simpletable.AlignCenter, Text: "Completed"},
		},
	}

	for _, row := range *todos {
		r := []*simpletable.Cell{
			{Text: strconv.Itoa(row.ID)},
			{Text: row.Title},
			{Text: strconv.FormatBool(row.Completed)},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleDefault)
	table.Println()
}

func AddTodo(title string) {
	todos := &models.Todo{}
	todo := todos.AddTodo(title, "todos.json")
	todos.Save("todos.json", todo)
	fmt.Println("Todo added successfully!")
	ListTodosCli()
}

func DeleteTodoById(id int) {
	todos := &models.Todo{}
	todos.Delete("todos.json", id)
	fmt.Printf("Todo '%d' deleted successfully!\n", id)
	ListTodosCli()
}
