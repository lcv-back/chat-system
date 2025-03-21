// Handle connections websocket and messages from client
package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Client là một client WebSocket
type Client struct {
	Connection *websocket.Conn
	Send       chan []byte
}

// Manager quản lý tất cả các client WebSocket
type Manager struct {
	Clients   map[*Client]bool
	Broadcast chan []byte
}

func NewManager() *Manager {
	return &Manager{
		Clients:   make(map[*Client]bool),
		Broadcast: make(chan []byte),
	}
}

// HandleConnections xử lý kết nối WebSocket từ client
func (manager *Manager) HandleConnections(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Cho phép mọi nguồn kết nối
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	client := &Client{Connection: conn, Send: make(chan []byte)}
	manager.Clients[client] = true
	fmt.Println("New client connected")

	go manager.HandleMessages(client)

	// Đọc tin nhắn từ client
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			delete(manager.Clients, client)
			close(client.Send)
			break
		}
		// Gửi tin nhắn đến tất cả các client kết nối
		manager.Broadcast <- message
	}
}

// HandleMessages gửi tin nhắn tới tất cả các client đã kết nối
func (manager *Manager) HandleMessages(client *Client) {
	for {
		message := <-manager.Broadcast
		for c := range manager.Clients {
			if c != client {
				c.Connection.WriteMessage(websocket.TextMessage, message)
			}
		}
	}
}
