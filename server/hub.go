package server

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
					delete(hub.Clients, client.ID)
					close(client.Send)
				}
			}
			hub.mu.Unlock()
		}

	}
}
