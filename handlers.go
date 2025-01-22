package main

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Message string `json:"message"`
}

func secureEndpoint(w http.ResponseWriter, r *http.Request) {
    response := Response{Message: "This is a secure endpoint"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}