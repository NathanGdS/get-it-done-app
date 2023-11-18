package handlers

import (
	"fmt"
	"strconv"

	"github.com/alexeyco/simpletable"
	"github.com/nathangds/get-it-done-app/models"
)

func ListTodosCli(fileName string) {
	var todos = &models.Todos

	if len(*todos) == 0 {
		todoReference := &models.Todo{}
		todoReference.Load(fileName)
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

func AddTodo(title string, fileName string) {
	todos := &models.Todo{}
	todo := todos.AddTodo(title, fileName)
	todos.Save(fileName, todo)
	fmt.Println("Todo added successfully!")
	ListTodosCli(fileName)
}

func DeleteTodoById(id int, fileName string) {
	todos := &models.Todo{}
	todos.Delete(fileName, id)
	fmt.Printf("Todo '%d' deleted successfully!\n", id)
	ListTodosCli(fileName)
}

func CompleteTodoById(id int, fileName string) {
	todos := &models.Todo{}
	todos.Complete(fileName, id)
	fmt.Printf("Todo '%d' completed successfully!\n", id)
	ListTodosCli(fileName)
}
