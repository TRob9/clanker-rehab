package concepts

// 054. Goroutine Basics
var Concept054 = Concept{
	Number:      54,
	ID:          "goroutine",
	Category:    "Concurrency",
	Name:        "54. Goroutine Basics",
	Description: "Launch goroutine with go",
	Instruction: "Launch an anonymous function in a goroutine that prints the word "hello", then sleep for 100 milliseconds to allow it to complete",
	Boilerplate: `package main

import (
	"fmt"
	"time"
)

func main() {
	// Your code here
}`,
	Answer: `package main

import (
	"fmt"
	"time"
)

func main() {
	go func() { fmt.Println("hello") }()
	time.Sleep(100 * time.Millisecond)
}`,
	ExpectedOutput: "hello",
	Difficulty:     "beginner",
	Explanation:    "Goroutines are lightweight threads managed by the Go runtime. The 'go' keyword launches a function in a new goroutine. Goroutines are cheap - you can have thousands. They run concurrently with the calling code.",
	Example:        "func sayHello() {\n    fmt.Println(\"hello\")\n}\n\ngo sayHello()  // runs concurrently\n\n// Anonymous function:\ngo func() {\n    fmt.Println(\"world\")\n}()\n\n// Multiple goroutines:\nfor i := 0; i < 10; i++ {\n    go process(i)\n}",
	UseCase:        "Use goroutines for concurrent tasks, I/O operations, parallel processing, and background work. Perfect for handling multiple connections, processing queues, or running periodic tasks. Goroutines are Go's killer feature.",
	Prerequisites:  []string{"function-value"},
	RelatedTopics:  []string{"channel-create", "waitgroup", "mutex"},
	DocsURL:        "https://go.dev/tour/concurrency/1",
}

func init() {
	Register(Concept054)
}
