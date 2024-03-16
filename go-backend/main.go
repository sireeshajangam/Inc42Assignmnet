// go-backend/main.go
//exam
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Text string `json:"message"`
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	message := Message{Text: "Hello from Go backend!"}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(message); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
