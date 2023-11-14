package models

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var Todos []Todo = []Todo{
	{ID: 1, Title: "Add Todos", Completed: false},
}

func (*Todo) ListTodos() []Todo {
	return Todos
}
