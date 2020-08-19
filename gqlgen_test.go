package main

import (
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/require"
	"github.com/ygorrodriguesdft/gqlgen-benchmark/gqlgen-server/generated"
	"github.com/ygorrodriguesdft/gqlgen-benchmark/gqlgen-server/model"
	"github.com/ygorrodriguesdft/gqlgen-benchmark/gqlgen-server/resolver"
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
		require.Equal(t, "T1", resp.Todos[0].ID)
	})
}

func BenchmarkGqlgenHello(b *testing.B) {
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}})))
	var resp struct {
		Hello string
	}
	c.MustPost(`query { hello }`, &resp)
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
