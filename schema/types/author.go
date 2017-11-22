package types

import "github.com/graphql-go/graphql"

var AuthorConfig = graphql.ObjectConfig{
	Name:        "Author",
	Description: "Authors write blog posts",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
	},
}
var Author = graphql.NewObject(AuthorConfig)
