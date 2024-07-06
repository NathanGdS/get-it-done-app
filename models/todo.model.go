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

func (t *Todo) LoadTodos(filename string) ([]Todo, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}

	if len(file) == 0 {
		return nil, err
	}
	err = json.Unmarshal(file, &Todos)
	if err != nil {
		return nil, err
	}

	return Todos, nil
}

func (t *Todo) AddTodo(title string, filename string) Todo {
	todos, err := t.LoadTodos(filename)

	if err != nil {
		log.Fatal(err)
	}

	actualSize := len(todos)
	if actualSize > 0 {
		actualSize = todos[actualSize-1].ID
	}
	newTodo := Todo{ID: actualSize + 1, Title: title, Completed: false}

	return newTodo
}

func (t *Todo) Save(filename string, todo Todo) {
	existentContent, err := t.Load(filename)

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

func (t *Todo) Load(filename string) (*[]Todo, error) {
	todos := &[]Todo{}

	file, err := os.OpenFile(filename, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return todos, nil
	}

	err = json.Unmarshal(data, &todos)

	if err != nil {
		return nil, err
	}

	Todos = *todos
	return todos, nil
}

func (t *Todo) Delete(filename string, id int) {
	existentContent, err := t.Load(filename)

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
	existentContent, err := t.Load(filename)

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
