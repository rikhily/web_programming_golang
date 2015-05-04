package main

import (
    "fmt"
    "net/http"
)

//takes the slice of bytes after the root
//and displays the text onto the localhost
// to test --> go to http://http://localhost:8080/<user_string>
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
