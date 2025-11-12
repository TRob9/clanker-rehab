package concepts

// 064. fmt.Println
var Concept064 = Concept{
	Number:      64,
	ID:          "fmt-println",
	Category:    "Standard Library",
	Name:        "64. fmt.Println",
	Description: "Print with newline",
	Instruction: "Print the word "hello" using the Println function from the fmt package",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	fmt.Println("hello")
}`,
	ExpectedOutput: "hello",
	Difficulty:     "beginner",
	Explanation:    "fmt.Println prints values to standard output, adding spaces between arguments and a newline at the end. It's the most common way to print output in Go. Values are formatted using their default format.",
	Example:        "fmt.Println(\"hello\")           // hello\\\nfmt.Println(\"a\", \"b\", \"c\")   // a b c\\\nfmt.Println(42)               // 42\\\nfmt.Println(true, false)      // true false",
	UseCase:        "Use fmt.Println for quick debugging, simple output, and logging. Perfect for development and learning. For production logging, consider structured logging libraries. For formatted output, use fmt.Printf instead.",
	Prerequisites:  nil,
	RelatedTopics:  []string{"fmt-printf"},
	DocsURL:        "https://pkg.go.dev/fmt#Println",
}

func init() {
	Register(Concept064)
}
