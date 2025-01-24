package main

import (
	"encoding/json"
	"net/http"
	"sync"
)

// Structure pour représenter un message
type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

// Exemple de données
var messages = []Message{
	{ID: 1, Content: "Hello, World!"},
	{ID: 2, Content: "Welcome to the Go API."},
}

var mu sync.Mutex

// Handler pour obtenir tous les messages
func getMessages(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

// Fonction principale
func main() {
	http.HandleFunc("/messages", getMessages)

	// Démarrer le serveur
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
