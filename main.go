package main

import (
    "fmt"
    "net/http"
		"os"
		"github.com/joho/godotenv"
)

func main() {
    godotenv.Load()

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello, Go backend is working!")
    })

    fmt.Println("Server running on port", port)
    http.ListenAndServe(":"+port, nil)
}
