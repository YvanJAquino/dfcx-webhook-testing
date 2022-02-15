package handlers

import (
	"encoding/json"
	"net/http"

	cx "github.com/yaq-cc/graffiti/godfcx"
)

type Message struct {
	Message string `json:"message,omitempty"`
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	msg := Message{
		Message: "Hello World!",
	}
	json.NewEncoder(w).Encode(&msg)

}

func TestingHandler(w http.ResponseWriter, r *http.Request) {
	var whresp cx.WebhookResponse
	whresp.TextResponse(w, "It works, Jim!")
}
