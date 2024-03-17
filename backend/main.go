package main    //example

import (
	"encoding/json"
	"log"
	"net/http"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func getPostsHandler(w http.ResponseWriter, r *http.Request) {
	// Simulate fetching posts from WordPress
	posts := []Post{
		{ID: 1, Title: "First Post", Body: "This is the body of the first post."},
		{ID: 2, Title: "Second Post", Body: "This is the body of the second post."},
	}

	// Convert posts to JSON
	jsonData, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, "Unable to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set response headers and write JSON data
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/posts", getPostsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
