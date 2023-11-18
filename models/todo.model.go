package models

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var Todos []Todo

func (t *Todo) LoadTodos(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, &Todos)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todo) AddTodo(title string, filename string) Todo {
	err := t.LoadTodos("todos.json")

	if err != nil {
		log.Fatal(err)
	}

	actualSize := len(Todos)
	newTodo := Todo{ID: actualSize + 1, Title: title, Completed: false}

	return newTodo
}

func (t *Todo) Save(filename string, todo Todo) {
	err, existentContent := t.Load(filename)

	if err != nil {
		log.Fatal(err)
	}

	newContent := append(*existentContent, todo)

	data, err := json.Marshal(newContent)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filename, data, 0644)

	if err != nil {
		log.Fatal(err)
	}

	Todos = newContent
}

func (t *Todo) Load(filename string) (error, *[]Todo) {
	todos := &[]Todo{}

	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return err, nil
	}
	defer file.Close()

	data, err := os.ReadFile(filename)

	if err != nil {
		return err, nil
	}

	if len(data) == 0 {
		return nil, todos
	}

	err = json.Unmarshal(data, &todos)

	if err != nil {
		return err, nil
	}

	Todos = *todos
	return nil, todos
}

func (t *Todo) Delete(filename string, id int) {
	err, existentContent := t.Load(filename)

	if err != nil {
		log.Fatal(err)
	}

	var newContent []Todo

	for _, todo := range *existentContent {
		if todo.ID != id {
			newContent = append(newContent, todo)
		}
	}

	if newContent == nil {
		newContent = []Todo{}
	}

	if len(newContent) == len(*existentContent) {
		log.Fatalf("Todo of id '%d' not found", id)
	}

	data, err := json.Marshal(newContent)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filename, data, 0644)

	if err != nil {
		log.Fatal(err)
	}

	Todos = newContent
}

func (t *Todo) Complete(filename string, id int) {
	err, existentContent := t.Load(filename)

	if err != nil {
		log.Fatal(err)
	}

	var newContent []Todo
	var founded bool = false

	for _, todo := range *existentContent {
		if todo.ID == id {
			todo.Completed = true
			founded = true
		}
		newContent = append(newContent, todo)
	}

	if !founded {
		log.Fatalf("Todo of id '%d' not found", id)
	}

	data, err := json.Marshal(newContent)

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(filename, data, 0644)

	if err != nil {
		log.Fatal(err)
	}

	Todos = newContent
}
