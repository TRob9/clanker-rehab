package concepts

// 061. Select with Default
var Concept061 = Concept{
	Number:      61,
	ID:          "select-default",
	Category:    "Concurrency",
	Name:        "61. Select with Default",
	Description: "Non-blocking select",
	Instruction: "Create an unbuffered channel of integers, then use a select statement with one case to receive from the channel and a default case that prints the word none",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	ch := make(chan int)
	select {
	case <-ch:
		fmt.Println("received")
	default:
		fmt.Println("none")
	}
}`,
	ExpectedOutput: "none",
	Difficulty:     "beginner",
	Explanation:    "Adding a default case to select makes it non-blocking. If no channel operation is ready, the default case executes immediately. This prevents select from waiting.",
	Example:        "ch := make(chan int)\n\nselect {\ncase v := <-ch:\n    fmt.Println(\"received\", v)\ndefault:\n    fmt.Println(\"no value ready\")  // executes immediately\n}\n\n// Try to send without blocking:\nselect {\ncase ch <- 42:\n    fmt.Println(\"sent\")\ndefault:\n    fmt.Println(\"channel blocked\")\n}",
	UseCase:        "Use default case for non-blocking channel operations, polling channels, avoiding deadlocks, or implementing try-send/try-receive. Useful in event loops and when you need to proceed regardless of channel state.",
	Prerequisites:  []string{"select"},
	RelatedTopics:  []string{"select", "select-timeout"},
	DocsURL:        "https://go.dev/tour/concurrency/6",
}

func init() {
	Register(Concept061)
}
