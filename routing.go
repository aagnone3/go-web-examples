package main

// More info at https://gowebexamples.com/routes-using-gorilla-mux/

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func launchMuxServer() {
    r := mux.NewRouter()

    r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, I see you've requested %s\n", r.URL.Path)
    })

    r.HandleFunc("/example_get", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, I see you've hit the %s endpoint. param: %s\n", r.URL.Path, r.URL.Query().Get("param"))
    })

    r.HandleFunc("/nested/{request}/page/{number}", func (w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        fmt.Fprintf(w, "Request: %s, page number: %s\n", vars["request"], vars["number"])
    })

    // define a file server for serving static assets
    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    fmt.Println("Listening on port 80...")
    http.ListenAndServe(":80", r)
}
