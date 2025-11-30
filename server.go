package main

import (
	"bufio"
	"fmt"
	"net"
	"sync"
)

type Client struct {
	id   int
	name string
	conn net.Conn
	ch   chan string
}

var (
	clients      = make(map[int]*Client)
	clientsMutex sync.Mutex
	broadcastCh  = make(chan string)
	nextID       = 1
)

func main() {
	go broadcaster()
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server running on port 9000")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}

		clientsMutex.Lock()
		id := nextID
		nextID++
		name := fmt.Sprintf("User%d", id)
		client := &Client{id: id, name: name, conn: conn, ch: make(chan string)}
		clients[id] = client
		clientsMutex.Unlock()

		// Broadcast join message to other clients
		broadcastCh <- fmt.Sprintf("%s joined", name)

		go handleClient(client)
		go sendMessages(client)
	}
}

func broadcaster() {
	for msg := range broadcastCh {
		fmt.Println(msg)
		clientsMutex.Lock()
		for _, c := range clients {
			// لمنع self-echo: لو الرسالة هي رسالة شخصية (UserX:) تجاهل العميل نفسه
			if len(msg) >= 5 && msg[:4] == c.name && msg[4] == ':' {
				continue
			}
			// لمنع العميل من رؤية join/leave الخاص به
			if (msg[len(msg)-6:] == "joined" || msg[len(msg)-4:] == "left") && msg[:len(c.name)] == c.name {
				continue
			}
			c.ch <- msg
		}
		clientsMutex.Unlock()
	}
}

func handleClient(c *Client) {
	reader := bufio.NewScanner(c.conn)
	for reader.Scan() {
		body := reader.Text()
		if body == "" {
			continue
		}
		msg := fmt.Sprintf("%s: %s", c.name, body)
		broadcastCh <- msg
	}
	clientsMutex.Lock()
	delete(clients, c.id)
	clientsMutex.Unlock()

	broadcastCh <- fmt.Sprintf("%s left", c.name)
	c.conn.Close()
}

func sendMessages(c *Client) {
	for msg := range c.ch {
		fmt.Fprintln(c.conn, msg)
	}
}
