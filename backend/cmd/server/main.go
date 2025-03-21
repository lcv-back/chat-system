package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Cho phép mọi kết nối
	},
}

// Client là một client WebSocket
type Client struct {
	Connection *websocket.Conn
	Send       chan []byte
	User       string
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

	// Lắng nghe tin nhắn từ client
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			delete(manager.Clients, client)
			close(client.Send)
			break
		}

		if messageType == websocket.TextMessage {
			// Xử lý tin nhắn dạng văn bản
			fmt.Println("Received a text message: ", string(message))
		} else if messageType == websocket.BinaryMessage {
			// Xử lý tin nhắn dạng nhị phân
			fmt.Println("Received a binary message")
		}

		// Kiểm tra xem đây là tin nhắn hay thông báo "typing"
		var msgData map[string]interface{}
		if err := json.Unmarshal(message, &msgData); err == nil {
			if msgData["type"] == "typing" {
				// Gửi thông báo "typing" đến tất cả client khác
				manager.Broadcast <- []byte(fmt.Sprintf(`{"type":"typing","user":"%s"}`, msgData["user"]))
			} else if msgData["type"] == "message" {
				// Gửi tin nhắn đến tất cả client khác
				manager.Broadcast <- []byte(fmt.Sprintf(`{"type":"message","message":"%s","user":"%s"}`, msgData["message"], msgData["user"]))
			}
		}
	}
}

// Broadcast tin nhắn đến tất cả client kết nối
func (manager *Manager) BroadcastMessages() {
	for {
		select {
		case message := <-manager.Broadcast:
			// Gửi tin nhắn đến tất cả các client
			for client := range manager.Clients {
				client.Connection.WriteMessage(websocket.TextMessage, message)
			}
		}
	}
}

func main() {
	manager := NewManager()

	// Khởi tạo server và WebSocket handler
	http.HandleFunc("/ws", manager.HandleConnections)

	// Chạy broadcast thread
	go manager.BroadcastMessages()

	// Cấu hình server
	serverAddr := ":1234"
	fmt.Println("Server started on", serverAddr)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		log.Fatal("Server failed: ", err)
	}
}
