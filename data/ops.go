package data

import (
	"fmt"
	"errors"
)

func GetPosts() []Post {
	return DataPosts
}

func GetPost(postID int) (Post, error) {
	for _, v := range DataPosts {
		if v.ID == postID {
			return v, nil
		}
	}
	return Post{}, errors.New("No record found")
}

func AddPost(p Post) Post {
	p.ID = len(DataPosts) + 1
	fmt.Println(p)
	DataPosts = append(DataPosts, p)

	return p
}

func GetComments(postID int) []Comment {
	var xc []Comment
	for _, v := range DataComments {
		if v.PostID == postID {
			xc = append(xc, v)
		}
	}
	return xc
}

func GetAuthors() []Author {
	return DataAuthors
}

func GetAuthor(authorID int) Author {
	for _, v := range DataAuthors {
		if v.ID == authorID {
			return v
		}
	}
	return Author{}
}

func GetAuthorPosts(authorID int) []Post {
	var xp []Post
	for _, v := range DataPosts {
		if v.AuthorID == authorID {
			xp = append(xp, v)
		}
	}
	return xp
}
