package server

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	checkOrigin: func(r *http.Request) bool {
		return true // allow all connections to connect not recommended in prod
	},
}
