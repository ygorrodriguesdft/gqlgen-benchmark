package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/ygorrodriguesdft/gqlgen-benchmark/graphql-go-server/schema/todo"
)

var queryFields = graphql.Fields{
	"listTodos": todo.ListTodos,
}

var mutationFields = graphql.Fields{
	"createTodo": todo.CreateTodo,
}

// defines the object config
var rootQuery = graphql.ObjectConfig{Name: "RootQuery", Fields: queryFields}
var rootMutation = graphql.ObjectConfig{Name: "RootMutation", Fields: mutationFields}

// defines a schema config
var schemaConfig = graphql.SchemaConfig{
	Query:    graphql.NewObject(rootQuery),
	Mutation: graphql.NewObject(rootMutation),
}

// Schema entire graphql schemas merged
var Schema, _ = graphql.NewSchema(schemaConfig)
