package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nyxoy77/websocket/auth"
	"github.com/nyxoy77/websocket/server"
)

func main() {
	// Initialize Redis

	// Create Gin router
	app := gin.Default()

	// Define WebSocket route

	app.GET("/ws", server.HandleWebSocketConnections)
	app.POST("/token", auth.HandleJWT)

	// Start WebSocket hub
	go server.RunHub()

	// Start server
	fmt.Println("WebSocket server running on ws://localhost:8080/ws")
	app.Run(":8080")
}
