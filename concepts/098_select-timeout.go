package concepts

// 098. Select with Timeout Pattern
var Concept098 = Concept{
	Number:      98,
	ID:          "select-timeout",
	Category:    "Concurrency",
	Name:        "98. Select with Timeout Pattern",
	Description: "Use select with time.After for timeout handling",
	Instruction: "Create a channel, use select to either receive from it or timeout after 10 milliseconds using the After function from the time package. Print the word timeout when the timeout case triggers.",
	Boilerplate: `package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// Your code here
}`,
	Answer: `package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	select {
	case <-ch:
		fmt.Println("received")
	case <-time.After(10 * time.Millisecond):
		fmt.Println("timeout")
	}
}`,
	ExpectedOutput: "timeout",
	Difficulty:     "advanced",
	Explanation:    "The select-timeout pattern combines select with time.After() to implement timeouts for channel operations. time.After() returns a channel that sends the current time after the specified duration, making it perfect for timeout cases.",
	Example:        "ch := make(chan string)\n\ngo func() {\n    time.Sleep(2 * time.Second)\n    ch <- \"result\"\n}()\n\nselect {\ncase msg := <-ch:\n    fmt.Println(\"received:\", msg)\ncase <-time.After(1 * time.Second):\n    fmt.Println(\"timeout!\")  // This executes\n}",
	UseCase:        "Use select with time.After for network requests, API calls, database queries, or any operation that shouldn't block indefinitely. Essential for building responsive systems that gracefully handle slow or unresponsive operations. Prevents goroutine leaks from hanging operations.",
	Prerequisites:  []string{"select", "channel-send-receive"},
	RelatedTopics:  []string{"select", "context-with-timeout", "select-default"},
	DocsURL:        "https://go.dev/blog/concurrency-timeouts",
}

func init() {
	Register(Concept098)
}
