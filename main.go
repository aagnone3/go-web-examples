package main

import (
    "net/http"
    "fmt"
)

const (
    USE_MUX = false
)

func launchNativeServer() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, I see you've requested %s\n", r.URL.Path)
    })

    // r.URL.path holds the endpoint
    // r.URL.Query() holds GET request parameters
    // r.FormValue(<param>) holds POST parameters
    http.HandleFunc("/example_get", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, I see you've hit the %s endpoint. param: %s\n", r.URL.Path, r.URL.Query().Get("param"))
    })

    // define a file server for serving static assets
    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    http.HandleFunc("/middleware", Chain(GreetRequest, ValidateRequestMethod("GET"), Logging()))

    fmt.Println("Listening on port 80...")
    http.ListenAndServe(":80", nil)
}

func main() {
    // db := dbConnection()
    // username := getUser(db, 1)
    // names := [4]string{"cmathis", "hagnone", "mlozano", "rlozano"}
    // for _, name := range(names) {
    //     createUser(db, name, "password")
    // }

    // users := getUsers(db)
    // for _, user := range(users) {
    //     fmt.Println(user)
    // }

    if USE_MUX == true {
        launchMuxServer()
    } else {
        launchNativeServer()
    }
}
