package main

import (
	"log"
	"net/http"

	"github.com/YvanJAquino/dfcx-webhook-testing/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/hello-dfcx", handlers.TestingHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
