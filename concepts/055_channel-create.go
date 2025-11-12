package concepts

// 055. Channel Creation
var Concept055 = Concept{
	Number:      55,
	ID:          "channel-create",
	Category:    "Concurrency",
	Name:        "55. Channel Creation",
	Description: "Create channel with make",
	Instruction: "Create a channel of integers using make, launch a goroutine that sends 42 to the channel, then receive from the channel and print the value",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() { ch <- 42 }()
	fmt.Println(<-ch)
}`,
	ExpectedOutput: "42",
	Difficulty:     "beginner",
	Explanation:    "Channels are typed conduits for communication between goroutines. Create channels with make(chan Type). Send and receive operations block by default, providing synchronization. Channels enable safe communication without explicit locks.",
	Example:        "ch := make(chan int)\n\ngo func() {\n    ch <- 42  // send to channel\n}()\n\nvalue := <-ch  // receive from channel\n\n// Channels synchronize:\ngo func() {\n    result := compute()\n    ch <- result  // blocks until received\n}()\nresult := <-ch  // blocks until sent",
	UseCase:        "Use channels to communicate between goroutines, synchronize execution, pass ownership of data, and distribute work. Channels are Go's primary synchronization mechanism. Prefer channels over shared memory + locks.",
	Prerequisites:  []string{"goroutine"},
	RelatedTopics:  []string{"goroutine", "buffered-channel", "channel-close"},
	DocsURL:        "https://go.dev/tour/concurrency/2",
}

func init() {
	Register(Concept055)
}
