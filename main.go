package main

import (
    "net/http"
    "log"
    "encoding/json"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        json.NewEncoder(writer).Encode(struct {
            Code int
        }{
            Code: 3000,
        })
    })
    log.Println("Start listen on 8080")
    http.ListenAndServe(":8080", mux)
}
