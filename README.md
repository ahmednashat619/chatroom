# chatroom
## 📂 Project Link  
[Google Drive Folder](https://drive.google.com/drive/folders/184gzPYJCiCA0_psf9jr-2ojrQHlvpdaN?usp=drive_link)

## 🧑‍💻 Overview  
This project implements a simple chatroom system using Go’s RPC mechanism. The system consists of:

- A **server** (`server.go`) that listens for client connections, stores all messages, and returns the full chat history each time a new message is sent.  
- A **client** (`client.go`) that connects to the server, sends messages, fetches and displays the chat history.  

## 🎯 Features  
- Clients run indefinitely until the user types `"exit"` or interrupts the program (Ctrl+C).  
- Each message from a client is stored by the server in a slice of messages.  
- After sending a message, the client receives the full chat history and displays it.  
- Thread-safe handling on the server is implemented (with a mutex) to allow multiple clients concurrently.
