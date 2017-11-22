package data

type Post struct {
	ID       int    `json:"id"`
	AuthorID int    `json:"authorId"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type Author struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Comment struct {
	ID      int    `json:"id"`
	PostID  int    `json:"postId"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

// dummy data
var DataPosts = []Post{
	{1, 1, "this is the way we roll", "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem"},
	{2, 1, "A different perspective", "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem"},
	{3, 2, "Go to town", "Lorem ipsum dolor sit amet, consectetuer adipiscing elit. Aenean commodo ligula eget dolor. Aenean massa. Cum sociis natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Donec quam felis, ultricies nec, pellentesque eu, pretium quis, sem"},
}

var DataAuthors = []Author{
	{1, "Mike Donnici", "michael@mesa.net.au"},
	{2, "Matt Donnici", "oi@mattdonnici.com"},
}

var DataComments = []Comment{
	{1, 1, "Anonymous", "This is a great article"},
	{2, 2, "Barry Anonymous", "Loved it!"},
	{3, 2, "Mary Anonymous", "A bit oo much to cope with!"},
	{4, 3, "John Smith", "Too many typos"},
	{5, 2, "Delta Fan Tasmo", "This could have been better"},
	{6, 1, "Rainy Dayz", "This is exactly what I needed"},
}
