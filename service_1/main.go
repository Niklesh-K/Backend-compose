package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from Golang Service 1!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting server on :5000...")
    http.ListenAndServe(":5000", nil)
}
