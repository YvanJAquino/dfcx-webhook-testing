package main

import (
	"log"
	"net/http"

	"github.com/YvanJAquino/dfcx-webhook-testing/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)
	err := http.ListenAndServe(":8083", nil)
	if err != nil {
		log.Fatal(err)
	}
}
