package main

import (
	"log"
	"net/http"
    "encoding/json"
    "os"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"
)

type Message struct {
	Response     string
	StatusCode   uint
}

var DefaultEndPoint = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var message Message

    message.Response = "Hello World!"
    message.StatusCode = 200
    jsonmessage, _ := json.Marshal(message)
    w.Write([]byte(jsonmessage))

})

var SecondEndPoint = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var message Message

    message.Response = "Yes, Scaling funds Is The best Out There. And i'm sure about that!!!!"
    message.StatusCode = 200
    jsonmessage, _ := json.Marshal(message)
    w.Write([]byte(jsonmessage))

})

func main() {
    router := mux.NewRouter()
    router.Handle("/", DefaultEndPoint).Methods("GET")
    router.Handle("/scalingfunds", SecondEndPoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, router)))
}


