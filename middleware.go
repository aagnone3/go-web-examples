package main

import (
    "log"
    "net/http"
    "time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc
type Handler func(w http.ResponseWriter, r *http.Request)

// Logging logs all requests and the elapsed time to process
func Logging() Middleware {

    // Create the middleware
    return func(f http.HandlerFunc) http.HandlerFunc {

        // Define the middleware's behavior
        return func(w http.ResponseWriter, r *http.Request) {
            start := time.Now()
            defer func() { log.Println(r.URL.Path, time.Since(start)) }()

            // Call the next middleware in the chain
            f(w, r)
        }
    }
}

// ValidateRequestMethod ensures that a request can only be made with the given method.
// If a request is made with a disallowed method, a 400 is returned.
func ValidateRequestMethod(methodRequired string) Middleware {

    // Create the middleware
    return func(f http.HandlerFunc) http.HandlerFunc {

        // Define the middleware's behavior
        return func(w http.ResponseWriter, r *http.Request) {

            if r.Method != methodRequired {
                http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
                return
            }

            // Call the next middleware in the chain
            f(w, r)
        }
    }
}

// Chain applied one or several middleware to http.HandlerFunc, returning its transform
// for repeated use
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
    for _, middleware := range(middlewares) {
        f = middleware(f)
    }
    return f
}

// GreetRequest represents the first processing of a request, prior to all middleware
func GreetRequest(w http.ResponseWriter, r *http.Request) {
    log.Printf("First hop of middleware request to %s\n", r.URL.Path)
}

// ProcessRequest represents the core processing of a request, after all middleware
func ProcessRequest(w http.ResponseWriter, r *http.Request) {
    // w.Fprintf(w, "First hop of middleware on request to %s\n", r.URL.Path)
}
