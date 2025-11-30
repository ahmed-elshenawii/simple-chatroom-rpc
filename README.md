# Simple Chatroom RPC

This repository contains a simple real-time chat application built using Go and net/rpc.
It includes:

* server.go — RPC chat server
* client.go — Chat client (CLI)

---

## 🎥 Demo Recording

[Click here to watch the demo](https://drive.google.com/file/d/1qqBQyplN591b0Ed2ivn6BkMZYM_6TuSX/view?usp=drive_link)

---

## Project Structure

```
/simple-chatroom-rpc
│-- server.go
│-- client.go
│-- README.md
```

---

##  How to Run the Project

### 1️⃣ Run the Server

```bash
go run server.go
```

### 2️⃣ Run the Client

```bash
go run client.go
```

---

##  How It Works

### Server

* Listens on `127.0.0.1:1234`
* Provides RPC method `ChatService.SendMessage`
* Stores chat messages in memory, with timestamp
* Returns full chat history to any client

### Client

* Connects to server via RPC
* User enters their name once
* Every message sent → server updates history → returns full chat history
* Type `exit` to quit

---

##  Example Chat Output

```
[12:33:10] Ahmed: Hello
[12:33:15] Sara: Hi Ahmed!
[12:33:20] Omar: Welcome guys
```

---

##  Notes

* Chat history is in-memory only — resets if server restarts.
* Works on any OS with Go installed.

---

##  Prepared by 

Ahmed Elshenawy
