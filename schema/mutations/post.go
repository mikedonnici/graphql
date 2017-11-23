package mutations

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/mikedonnici/graphql/schema/types"
	"github.com/mikedonnici/graphql/data"
)

var AddPost = &graphql.Field{
	Name:        "AddPost",
	Description: "Add a post",
	Type:        types.Post,
	Args: graphql.FieldConfigArgument{
		"post": &graphql.ArgumentConfig{
			Type:        types.PostInput,
			Description: "A post object",
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		// post object is received as a map[string]interface{}
		// need to 'convert' to a data.Post for AddPost()
		// Don't need validation as GraphQL will not allow incorrect types
		post, ok := p.Args["post"].(map[string]interface{})
		if ok {
			dp := data.Post{}
			dp.Title = post["title"].(string)
			dp.Content = post["content"].(string)
			dp.AuthorID = post["authorId"].(int)
			return data.AddPost(dp), nil
		}
		return nil, nil
	},
}
