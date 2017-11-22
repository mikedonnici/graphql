package types

import "github.com/graphql-go/graphql"

var Comment = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Comment",
	Description: "Comments made on posts",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var CommentInput = graphql.NewObject(graphql.ObjectConfig{
	Name:        "CommentInput",
	Description: "Type to add a comment",
	Fields: graphql.Fields{
		"postId": &graphql.Field{
			Type: &graphql.NonNull{OfType: graphql.Int},
		},
		"name": &graphql.Field{
			Type: &graphql.NonNull{OfType: graphql.String},
		},
		"content": &graphql.Field{
			Type: &graphql.NonNull{OfType: graphql.String},
		},
	},
})
