package concepts

// 060. Select Statement
var Concept060 = Concept{
	Number:      60,
	ID:          "select",
	Category:    "Concurrency",
	Name:        "60. Select Statement",
	Description: "Select on multiple channels",
	Instruction: "Create a buffered channel of integers with capacity 1, send 5 to it, then use a select statement with a single case that receives the value and prints it",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 5
	select {
	case v := <-ch:
		fmt.Println(v)
	}
}`,
	ExpectedOutput: "5",
	Difficulty:     "beginner",
	Explanation:    "Select waits on multiple channel operations. It blocks until one case can proceed, then executes that case. If multiple cases are ready, one is chosen at random. Select is like switch but for channels.",
	Example:        "ch1 := make(chan string)\nch2 := make(chan string)\n\ngo func() { ch1 <- \"one\" }()\ngo func() { ch2 <- \"two\" }()\n\nselect {\ncase msg1 := <-ch1:\n    fmt.Println(msg1)\ncase msg2 := <-ch2:\n    fmt.Println(msg2)\n}\n// Prints whichever arrives first",
	UseCase:        "Use select to wait on multiple channels, implement timeouts, handle multiple communication paths, or build non-blocking operations. Essential for complex concurrent patterns and multiplexing channels.",
	Prerequisites:  []string{"channel-send-receive"},
	RelatedTopics:  []string{"select-default", "select-timeout", "channel-send-receive"},
	DocsURL:        "https://go.dev/tour/concurrency/5",
}

func init() {
	Register(Concept060)
}
