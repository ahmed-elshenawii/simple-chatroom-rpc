# 💬 Simple RPC Chatroom in Go

## 📖 Description
This project implements a simple chatroom using Go’s **net/rpc** package.

- The server stores all messages in a shared history.  
- Each client can send messages and fetch the entire chat history from the server.  
- Multiple clients can connect and see all previous messages.  

---

## 🎥 Demo Video
Watch the demo here:  
[▶️ Click to watch](https://drive.google.com/file/d/1jcmgjZNCj_TvRvQ_7QeJJriOBHeNXwie/view?usp=drive_link)

---

## ⚙️ How to Run

### 🧩 Prerequisites
Make sure you have:
- **Go** (version 1.18 or higher)
- A **terminal** or **command prompt**

---

### 🚀 Steps to Run

Clone the repository
```bash
   git clone https://github.com/yourusername/simple-rpc-chatroom.git
   cd simple-rpc-chatroom

Run the server
   go run server.go

Open another terminal and run the client
   go run client.go

Start chatting
   Enter your name and start chatting!
   Type your message and press Enter.
   Type exit to leave the chat.

Example Output
   Enter your name: Ahmed
   > Hello everyone!
   ----- Chat History -----
   [01:41:32] Ahmed: Hello everyone!

   > Hi
   ----- Chat History -----
   [01:41:32] Ahmed: Hello everyone!
   [01:41:44] Ahmed: Hi

Error Handling
   Displays "Connection error:" if the server is down.
   Gracefully exits when typing "exit".

Project Structure
simple-rpc-chatroom/
│
├── server.go      # RPC server – stores and returns messages
├── client.go      # RPC client – sends messages and displays history
└── README.md      # Documentation file

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
