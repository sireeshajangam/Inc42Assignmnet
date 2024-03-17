// backend/main.go 
//ex

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/posts", GetPostsHandler)
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	// Handler logic goes here
	fmt.Fprintf(w, "First Post\nSecond Post")
}
