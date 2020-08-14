package todo

import (
	"github.com/graphql-go/graphql"
)

var TodoList []Todo

type Todo struct {
	ID    int
	Title string
	Done  bool
	User  User
}

type User struct {
	ID   int
	Name string
}

var ListTodos = &graphql.Field{
	Type:        graphql.NewList(TypeTodo),
	Description: "Get Full list of todos",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return TodoList, nil
	},
}

var CreateTodo = &graphql.Field{
	Type:        TypeTodo,
	Description: "Create new todo",
	Args: graphql.FieldConfigArgument{
		"text": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		text, _ := params.Args["text"].(string)
		newTodo := Todo{
			ID:    1,
			Title: text,
			Done:  false,
			User:  User{ID: 1, Name: "Oba"},
		}
		TodoList = append(TodoList, newTodo)
		return newTodo, nil
	},
}
