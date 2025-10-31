package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

// Message struct (same as on server)
type Message struct {
	Sender    string
	Content   string
	Timestamp string
}

func main() {
	// Try to connect to the server
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("❌ Connection error:", err)
	}
	defer client.Close()
	fmt.Println("✅ Connected to chat server")

	// Ask for user name
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("👋 Exiting chat...")
			break
		}

		// Prepare message
		msg := Message{Sender: name, Content: text}

		// Send message to server and get chat history
		var history []Message
		err := client.Call("ChatService.SendMessage", msg, &history)
		if err != nil {
			fmt.Println("❌ RPC Error:", err)
			break
		}

		// Display full chat history
		fmt.Println("----- Chat History -----")
		for _, m := range history {
			fmt.Printf("[%s] %s: %s\n", m.Timestamp, m.Sender, m.Content)
		}
		fmt.Println("------------------------")
	}
}
