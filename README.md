# graphql boilerplate

Playing with go-graphql...

**Todo**

- [x] Fake datastore
- [x] Basic query
- [x] Web handler (used graphql-handler) 
- [x] Get a basic mutation to work
- [x] Mutation with an 'object'?


### Objects as Input Types

After some confusion -> [the answer](https://stackoverflow.com/questions/40360936/graphql-go-use-an-object-as-input-argument-to-a-query)

So, to do a mutation that looks like this:

```graphql
mutation{
  addPost(
    post:{
      title: "The week I learned GraphQL",
      content: "This week I started to learn GraphQL..."
      authorId: 1
    }
  ){
    id
    title
  }
}
```

The PostInput type is:

```golang
var PostInput = graphql.NewInputObject(graphql.InputObjectConfig{
	Name:        "PostInput",
	Description: "Add new post type",
	Fields: graphql.InputObjectConfigFieldMap{
		"title": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.String},
		},
		"content": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.String},
		},
		"authorId": &graphql.InputObjectFieldConfig{
			Type: &graphql.NonNull{OfType: graphql.Int},
		},
	},
})
```

The `AddPost` Mutation field value is:

```golang
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
		// GraphQL handles validation
		post, ok := p.Args["post"].(map[string]interface{})
		fmt.Println(post)
		if ok {
			dp := data.Post{}
			dp.Title = post["title"].(string)
			dp.Content = post["content"].(string)
			dp.AuthorID = post["authorId"].(int)
			fmt.Println(dp)
			return data.AddPost(dp), nil
		}
		return nil, nil
	},
}
``` 

The tricky bit was that the arg `(post:{...})` end up as a `map[string]interface{}` and does not automagically become a `data.Post` - unless I have missed something.

So this line does the assertion and then can create a fresh data.Post value to send to data management land.

```golang
post, ok := p.Args["post"].(map[string]interface{})
```





 


