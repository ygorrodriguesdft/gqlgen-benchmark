package main

import (
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/ygorrodriguesdft/gqlgen-benchmark/graphql-go-server/schema"
)

func BenchmarkGraphqlGOHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		graphql.Do(graphql.Params{
			Schema:        schema.Schema,
			RequestString: `query { hello }`,
		})
	}
}

func BenchmarkGraphqlGOCreateTodo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		graphql.Do(graphql.Params{
			Schema:        schema.Schema,
			RequestString: `mutation {createTodo(text:"teste") {ID Title Done User{ID Name}}}`,
		})
	}
}

func BenchmarkGraphqlGOListTodos(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		graphql.Do(graphql.Params{
			Schema:        schema.Schema,
			RequestString: `query {listTodos {ID Title Done User{ID Name}}}`,
		})
	}
}
