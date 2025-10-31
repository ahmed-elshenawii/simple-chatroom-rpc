# 💬 Simple RPC Chatroom in Go

## 📘 Overview
A simple chatroom built using Go’s **net/rpc** package.  
Each client connects to the server, sends messages, and receives the full chat history.

---

## 🎥 Demo Video
Watch the running demo here:  
[▶️ Click to watch](https://drive.google.com/file/d/1qqBQyplN591b0Ed2ivn6BkMZYM_6TuSX/view?usp=drive_link)

---

## ✨ Features
💠 **RPC-based Communication** – Implements Go’s `net/rpc` for client-server interaction.  
💠 **Shared Chat History** – Every connected client sees the entire chat log.  
💠 **Multiple Clients Support** – Handles concurrent users seamlessly.  
💠 **Real-Time Updates** – Instantly displays all new messages for all clients.  
💠 **In-Memory Message Storage** – Keeps messages during the active session.  
💠 **Error Handling** – Displays clear messages when the server is unreachable.  
💠 **Graceful Exit** – Type `exit` to leave the chat safely.  

---

## 🧠 Example Output
Enter your name: Ahmed.

Welcome One! You've joined the chat. Type a message to see the chat history.

Enter message (or 'exit' to quit):

Hi

----- Chat History -----

[01:42:31] One: Hi

exit

👋 Goodbye!


---

## 👨‍💻 Prepared By
  ###   🌟 **Ahmed Ibrahim Ahmed Elshenawy**
  🎓   Faculty of Engineering – Department of Artificial Intelligence  
  📅   **October 2025**  
  💼   *Project: Simple RPC Chatroom using Go*  
  💬   *"Clean code. Clear communication."*
