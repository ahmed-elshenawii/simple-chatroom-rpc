package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
	"time"
)

// Message structure to store sender and content
type Message struct {
	Sender    string
	Content   string
	Timestamp string
}

// ChatService holds the chat history
type ChatService struct {
	mu      sync.Mutex
	history []Message
}

// SendMessage is called remotely by clients to send a message
func (c *ChatService) SendMessage(msg Message, reply *[]Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Add timestamp and store message
	msg.Timestamp = time.Now().Format("15:04:05")
	c.history = append(c.history, msg)

	// Return full chat history to the client
	*reply = c.history
	return nil
}

// GetHistory returns all stored messages (if client just wants to fetch)
func (c *ChatService) GetHistory(_ struct{}, reply *[]Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	*reply = c.history
	return nil
}

func main() {
	chatService := new(ChatService)
	rpc.Register(chatService)

	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Chat server is running on 127.0.0.1:1234...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
