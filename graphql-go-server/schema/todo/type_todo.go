package todo

import (
	"github.com/graphql-go/graphql"
)

// TypeTodo author data type
var TypeTodo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Todo",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.ID,
			},
			"Title": &graphql.Field{
				Type: graphql.String,
			},
			"Done": &graphql.Field{
				Type: graphql.Boolean,
			},
			"User": &graphql.Field{
				Type: TypeUser,
			},
		},
	},
)
