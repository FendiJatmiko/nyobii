package main

import (
	"net/http"

	"github.com/fendijatmiko/nyobiid"
	"golang.org/x/net/websocket"
)

func main() {

	http.Handle("/nyobi", websocket.Handler(func(ws *websocket.Conn) {
		ws.Write([]byte(nyobiid.Nyobii()))
	}))
	http.ListenAndServe(":8080", nil)
}
