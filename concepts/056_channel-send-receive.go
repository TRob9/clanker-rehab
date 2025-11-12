package concepts

// 056. Channel Send/Receive
var Concept056 = Concept{
	Number:      56,
	ID:          "channel-send-receive",
	Category:    "Concurrency",
	Name:        "56. Channel Send/Receive",
	Description: "Send and receive from channel",
	Instruction: "Create a channel of strings, launch a goroutine that sends the message msg to the channel, receive the value into a variable, then print the variable",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() { ch <- "msg" }()
	x := <-ch
	fmt.Println(x)
}`,
	ExpectedOutput: "msg",
	Difficulty:     "beginner",
	Explanation:    "Send to a channel with ch <- value (arrow points into channel). Receive from a channel with value := <-ch (arrow points out). Both operations block until the other side is ready, providing automatic synchronization.",
	Example:        "ch := make(chan string)\n\n// Send (blocks until received):\ngo func() {\n    ch <- \"hello\"\n    ch <- \"world\"\n}()\n\n// Receive (blocks until sent):\nmsg1 := <-ch  // \"hello\"\nmsg2 := <-ch  // \"world\"\n\n// Discard value:\n<-ch  // receive but ignore",
	UseCase:        "Use send/receive for goroutine communication, passing results, signaling completion, or streaming data. The blocking behavior naturally synchronizes goroutines without explicit locks or condition variables.",
	Prerequisites:  []string{"channel-create"},
	RelatedTopics:  []string{"channel-create", "buffered-channel", "select"},
	DocsURL:        "https://go.dev/tour/concurrency/2",
}

func init() {
	Register(Concept056)
}
