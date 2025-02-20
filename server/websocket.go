package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"

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
	mu         sync.Mutex
}

var hub = Hub{
	Clients:    make(map[string]*Client),
	Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
}

func HandleWebSocketConnections(c *gin.Context) {
	// token := c.Query("token")
	// userID, err := auth.ValidateJWT(token)
	// if err != nil || userID == "" {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"error": "Invalid token",
	// 	})
	// }
	userID := c.Query("userID")
	if userID == "" {
		c.String(http.StatusUnauthorized, "Missing userID")
		return
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Error upgrading the websocket ", err)
		return
	}
	fmt.Println(userID)
	client := &Client{ID: userID, Conn: conn, Send: make(chan []byte)}
	hub.Register <- client
	go ReadMessages(client)
	go WriteMessages(client)
}

func ReadMessages(client *Client) {
	defer func() {
		hub.Unregister <- client // remove the client when disconnected
		client.Conn.Close()
	}()

	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		hub.Broadcast <- msg
	}
}

func WriteMessages(client *Client) {
	defer client.Conn.Close()

	for msg := range client.Send {
		err := client.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}
