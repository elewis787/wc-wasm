package main 

import (
    "log"
    "net/http"
)

func main() {
    fs := http.FileServer(http.Dir("./")) // Serve files from the current directory
    http.Handle("/", fs)

    log.Println("Listening on :8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}