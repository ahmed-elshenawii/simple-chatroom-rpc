package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	// Connect to server
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer client.Close()

	// Get user name
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Welcome %s! You've joined the chat. Type a message to see the chat history.\n", name)
	fmt.Println("Enter message (or 'exit' to quit):")

	for {
		fmt.Print("> ")
		message, _ := reader.ReadString('\n')
		message = strings.TrimSpace(message)

		if strings.ToLower(message) == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		var reply []string
		err = client.Call("ChatService.SendMessage", map[string]string{
			"name":    name,
			"message": message,
		}, &reply)

		if err != nil {
			fmt.Println("RPC Error:", err)
			break
		}

		fmt.Println("----- Chat History -----")
		for _, line := range reply {
			fmt.Println(line)
		}
	}
}
