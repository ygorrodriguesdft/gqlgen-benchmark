package todo

import "github.com/graphql-go/graphql"

// TypeUser author data type
var TypeUser = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.ID,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)
