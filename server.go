package main

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

// Message holds a single chat message
type Message struct {
	Sender  string
	Content string
}

// ChatServer represents the chatroom server
type ChatServer struct {
	mu       sync.Mutex
	messages []Message
}

// Args is used for sending messages from the client
type Args struct {
	Sender  string
	Content string
}

// SendMessage adds a message to the chat and returns all messages
func (c *ChatServer) SendMessage(args Args, reply *[]Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.messages = append(c.messages, Message{Sender: args.Sender, Content: args.Content})
	*reply = c.messages
	return nil
}

// GetMessages returns all chat history
func (c *ChatServer) GetMessages(dummy int, reply *[]Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	*reply = c.messages
	return nil
}

func main() {
	server := new(ChatServer)
	rpc.Register(server)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Chat server running on port 1234...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn)
	}
}
