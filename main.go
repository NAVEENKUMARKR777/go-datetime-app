package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

func dateTimeHandler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format("2006-01-02 15:04:05 MST")
    fmt.Fprintf(w, "<h1>Current Date & Time</h1>")
    fmt.Fprintf(w, "<p>%s</p>", currentTime)
}

func main() {
    http.HandleFunc("/", dateTimeHandler)
    fmt.Println("Starting server on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
