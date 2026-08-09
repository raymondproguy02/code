package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type Health struct {
    Status    string `json:"status"`
    Timestamp string `json:"timestamp"`
    Version   string `json:"version"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Health{
        Status:    "ok",
        Timestamp: time.Now().UTC().Format(time.RFC3339),
        Version:   "1.0.0",
    })
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Query().Get("name"))
}

func main() {
    http.HandleFunc("/health", healthHandler)
    http.HandleFunc("/hello", helloHandler)
    
    fmt.Println("🚀 Server starting on :8080")
    fmt.Println("📊 Health: http://localhost:8080/health")
    fmt.Println("👋 Hello: http://localhost:8080/hello?name=Go")
    
    http.ListenAndServe(":8080", nil)
}
