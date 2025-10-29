package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"strings"
)

type Message struct {
	Sender  string
	Content string
}

type Args struct {
	Sender  string
	Content string
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Welcome %s! You've joined the chat.\n", name)
	fmt.Println("Type a message to see the chat history.")

	for {
		fmt.Print("Enter message (or 'exit' to quit): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if strings.ToLower(text) == "exit" {
			fmt.Println("Exiting chat...")
			break
		}

		var chatHistory []Message
		args := Args{Sender: name, Content: text}

		err = client.Call("ChatServer.SendMessage", args, &chatHistory)
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}

		fmt.Println("\n--- Chat History ---")
		for _, msg := range chatHistory {
			fmt.Printf("%s: %s\n", msg.Sender, msg.Content)
		}
		fmt.Println()
	}
}
