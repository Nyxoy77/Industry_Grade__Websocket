package server

import "log"

func RunHub() {
	for {
		select {
		case client := <-hub.Register:
			hub.mu.Lock()
			hub.Clients[client.ID] = client
			hub.mu.Unlock()
		case client := <-hub.Unregister:
			hub.mu.Lock()
			delete(hub.Clients, client.ID)
			close(client.Send)
			hub.mu.Unlock()
		case message := <-hub.Broadcast:
			hub.mu.Lock()
			for _, client := range hub.Clients {
				select {
				case client.Send <- message:
				default:
					log.Println("Socket unresponsive ", client.ID, "Approval for removing")
					go func(c *Client) {
						hub.mu.Lock()
						delete(hub.Clients, client.ID)
						close(client.Send)
						hub.mu.Unlock()
					}(client)
				}
			}
			hub.mu.Unlock()
		}

	}
}
