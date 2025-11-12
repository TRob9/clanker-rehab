package concepts

// 093. Context with Timeout
var Concept093 = Concept{
	Number:      93,
	ID:          "context-with-timeout",
	Category:    "Concurrency",
	Name:        "93. Context with Timeout",
	Description: "Use context for cancellation and timeouts",
	Instruction: "Create a context with a 100 millisecond timeout using the WithTimeout function from the context package. Print the word "done" immediately without waiting for the timeout.",
	Boilerplate: `package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Your code here
}`,
	Answer: `package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	fmt.Println("done")
}`,
	ExpectedOutput: "done",
	Difficulty:     "advanced",
	Explanation:    "Context carries deadlines, cancellation signals, and request-scoped values across API boundaries. context.WithTimeout creates a context that automatically cancels after a duration. Check ctx.Done() channel or ctx.Err() to detect cancellation.",
	Example:        "ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)\ndefer cancel()  // always call cancel\n\nselect {\ncase <-time.After(1 * time.Second):\n    fmt.Println(\"work done\")\ncase <-ctx.Done():\n    fmt.Println(\"timeout:\", ctx.Err())\n}",
	UseCase:        "Use contexts for request timeouts, cancellation propagation, API calls with deadlines, and coordinating goroutine shutdown. Pass context as first parameter to functions. Essential for servers and concurrent operations.",
	Prerequisites:  []string{"goroutine", "select"},
	RelatedTopics:  []string{"select-timeout", "goroutine"},
	DocsURL:        "https://pkg.go.dev/context",
}

func init() {
	Register(Concept093)
}
