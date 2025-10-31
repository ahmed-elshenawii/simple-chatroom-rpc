# 💬 Simple RPC Chatroom in Go

## 📘 Project Description
This project is a simple **chatroom application** built using **Go’s `net/rpc` package**.  
It demonstrates **client-server communication** via Remote Procedure Calls (RPC).

### 🖥️ Server
- Stores all messages in a shared message list.  
- Returns the **entire chat history** to clients upon request.  
- Runs continuously, waiting for client connections.

### 💻 Client
- Connects to the RPC server.  
- Sends messages remotely through the server’s exposed procedure.  
- Fetches and displays the full chat history after each message.  
- Keeps running until the user types `"exit"` or presses `Ctrl + C`.

---

## 🎥 Demo Video
📺 Watch the running application here:  
👉 [**Click to Watch Demo**](https://drive.google.com/file/d/1jcmgjZNCj_TvRvQ_7QeJJriOBHeNXwie/view?usp=drive_link)

---

## ⚙️ How to Run

### 🧩 Prerequisites
Make sure you have:
- Go (version 1.18 or higher)
- A terminal or command prompt

### 🚀 Steps to Run

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/simple-rpc-chatroom.git
   cd simple-rpc-chatroom
   
2. Run the server
   go run server.go
   
3.Open another terminal and run the client
   go run client.go

4.Enter your name and start chatting!
   Type your message and press Enter.
   Type exit to leave the chat.

Example Output
   Enter your name: Ahmed
   > Hello everyone!
   ----- Chat History -----
   [01:41:32] Ahmed: Hello everyone!
   ------------------------
   > Hi
   ----- Chat History -----
   [01:41:32] Ahmed: Hello everyone!
   [01:41:44] Ahmed: Hi
   ------------------------

Error Handling
   Displays "Connection error:" if the server is down.
   Gracefully exits when typing "exit".

Project Structure
   simple-rpc-chatroom/
│
├── server.go        # RPC server – stores and returns messages
├── client.go        # RPC client – sends messages and displays history
└── README.md        # Documentation file

Documentation Summary
   This project demonstrates:
   RPC communication in Go
   Persistent message storage on the server
   Concurrent access by multiple clients
   Basic error handling for lost connections

Prepared by
   Ahmed Elshenawy
   Faculty of Engineering – Department of Artificial Intelligence
   October 2025
