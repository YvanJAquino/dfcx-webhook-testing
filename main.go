package main

import (
	"log"
	"net/http"

	"github.com/YvanJAquino/dfcx-webhook-testing/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.HelloHandler)
	http.HandleFunc("/hello-dfcx", handlers.TestingHandler)
	// err := http.ListenAndServe(":8083", nil)
	err := http.ListenAndServeTLS(":8443", "certFile.pem", "keyFile.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
