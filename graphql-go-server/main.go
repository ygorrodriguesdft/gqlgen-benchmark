package main

import (
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/ygorrodriguesdft/gqlgen-benchmark/graphql-go-server/schema"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main() {
	h := handler.New(&handler.Config{
		Schema:     &schema.Schema,
		Pretty:     true,
		GraphiQL:   false,
		Playground: true,
	})

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	fmt.Println("Server online na porta 8000")
	http.Handle("/graphql", h)
	http.ListenAndServe(":8000", nil)
}
