package data

// Data model structs
type Post struct {
	Id   string `json:"id"`
	Text string `json:"text"`
}

// Mock data
var latestPost = &Post{
	Id: "1",
	Text: "Hello World",
}

// Data getters/setters
func GetLatestPost() *Post {
	return latestPost
}