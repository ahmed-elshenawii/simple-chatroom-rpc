package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
	"time"
)

type ChatService struct {
	messages []string
	mu       sync.Mutex
}

func (c *ChatService) SendMessage(data map[string]string, reply *[]string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	name := data["name"]
	message := data["message"]
	timestamp := time.Now().Format("15:04:05")

	formatted := fmt.Sprintf("[%s] %s: %s", timestamp, name, message)
	c.messages = append(c.messages, formatted)

	*reply = append([]string(nil), c.messages...)
	fmt.Printf("%s\n", formatted)
	return nil
}

func main() {
	server := new(ChatService)
	rpc.Register(server)

	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		panic(err)
	}
	fmt.Println("Chat server running on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(conn)
	}
}
