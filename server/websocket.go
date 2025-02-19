package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

//Upgrader is a struct we cant use its functions without it being stored in an instance

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // allow all connections to connect not recommended in prod
	},
}

type Client struct {
	ID   string
	Conn *websocket.Conn
	Send chan []byte
}

type Hub struct {
	Clients    map[string]*Client
	Broadcast  chan []byte
	Register   chan *Client
	Unregister chan *Client
}

var hub = Hub{
	Clients:    make(map[string]*Client),
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
}

func HandleWebSocketConnections(c *gin.Context) {
	// token := c.Query("token")
	// userID,err := ValidateJWT(token)
}
