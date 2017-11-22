package types

import (
	"github.com/graphql-go/graphql"
	"github.com/mikedonnici/graphql/data"
)

var Post = graphql.NewObject(graphql.ObjectConfig{
	Name:        "Post",
	Description: "Blog post...",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"author": &graphql.Field{
			Type: Author,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// get the author id from the source of this data, which is data.Post.AuthorID
				// p.Source returns an interface{} so have to assert the value first :)
				authorID := p.Source.(data.Post).AuthorID
				return data.GetAuthor(authorID), nil
			},
		},
		"comments": &graphql.Field{
			Type:        graphql.NewList(Comment),
			Description: "Comments belonging to a post",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				postID := p.Source.(data.Post).ID
				return data.GetComments(postID), nil
			},
		},
	},
})

var PostInput = graphql.NewObject(graphql.ObjectConfig{
	Name:        "PostInput",
	Description: "Add new post type",
	Fields: graphql.Fields{
		"title": &graphql.Field{
			Type: &graphql.NonNull{OfType: graphql.Int},
		},
		"content": &graphql.Field{
			Type: &graphql.NonNull{OfType: graphql.String},
		},
		"authorId": &graphql.Field{
			Type: &graphql.NonNull{OfType: graphql.Int},
		},
	},
})
