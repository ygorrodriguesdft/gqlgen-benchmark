package main

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/graphql-go/graphql"
	"github.com/ygorrodriguesdft/gqlgen-benchmark/gqlgen-server/generated"
	"github.com/ygorrodriguesdft/gqlgen-benchmark/gqlgen-server/model"
	"github.com/ygorrodriguesdft/gqlgen-benchmark/gqlgen-server/resolver"
	"github.com/ygorrodriguesdft/gqlgen-benchmark/graphql-go-server/schema"
)

func TestTodo(t *testing.T) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}})))

	var resp struct {
		CreateTodo struct{ ID string }
	}

	c.MustPost(`mutation { createTodo(input:{text:"Fery important", userId:"1"}) { id } }`, &resp)

	t.Run("list todos", func(t *testing.T) {
		var resp struct {
			Todos []struct {
				Text string
				ID   string
			}
		}
		c.MustPost(`{
			todos{ id }
		}`, &resp)
		// require.Equal(t, "1", resp.Todo.ID)
	})
}

func BenchmarkGqlgenCreateTodo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}})))
		var resp struct {
			CreateTodo struct {
				ID   string
				Text string
				Done bool
				User model.User
			}
		}
		c.MustPost(`mutation { createTodo(input:{text:"Fery important", userId:"1"}) { id text done user{id name}} }`, &resp)
	}
}

func BenchmarkGqlgenListTodos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}})))
		var resp struct {
			Todos []struct {
				Text string
				ID   string
				Done bool
				User model.User
			}
		}
		c.MustPost(`{ todos{ id text done user{id name}} }`, &resp)
	}
}

func BenchmarkGraphqlGOCreateTodo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		graphql.Do(graphql.Params{
			Schema:        schema.Schema,
			RequestString: `mutation {createTodo(text:"teste") {ID Title Done User{ID Name}}}`,
		})
	}
}

func BenchmarkGraphqlGOListTodos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		graphql.Do(graphql.Params{
			Schema:        schema.Schema,
			RequestString: `query {listTodos {ID Title Done User{ID Name}}}`,
		})
	}
}
