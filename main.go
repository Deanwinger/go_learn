package main

import (
    "fmt"
	"net/http"
    "strings"
    "log"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    fmt.Println(r.Form)
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!")
}

func main() {
	http.HandleFunc("/", sayHelloName)
    err := http.ListenAndServe(":8000", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}