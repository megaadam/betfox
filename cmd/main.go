package main

import (
    "fmt"
    "log"
    "net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "root: %s!", r.URL.Path[1:])
}
func subHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "sub: %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", rootHandler)
    http.HandleFunc("/sub", rootHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

