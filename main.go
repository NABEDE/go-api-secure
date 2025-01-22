package main

import (
	"encoding/json"
	"net/http"
)

// Définir une structure pour les données
type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

// Exemple de données
var messages = []Message{
	{ID: 1, Content: "Hello, World!"},
	{ID: 2, Content: "Welcome to the Go API."},
}

// Handler pour obtenir tous les messages
func getMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

// Fonction principale
func main() {
	http.HandleFunc("/messages", getMessages)

	// Démarrer le serveur
	http.ListenAndServe(":8080", nil)
}