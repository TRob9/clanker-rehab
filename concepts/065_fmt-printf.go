package concepts

// 065. fmt.Printf
var Concept065 = Concept{
	Number:      65,
	ID:          "fmt-printf",
	Category:    "Standard Library",
	Name:        "65. fmt.Printf",
	Description: "Formatted print",
	Instruction: "Print the text "num: 5" using Printf from the fmt package with the integer format specifier",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	fmt.Printf("num: %d", 5)
}`,
	ExpectedOutput: "num: 5",
	Difficulty:     "beginner",
	Explanation:    "fmt.Printf prints formatted output using format specifiers (verbs). Common verbs: %d (int), %s (string), %f (float), %v (any value), %T (type). It doesn't add a newline unless you include \\\n in the format string.",
	Example:        "fmt.Printf(\"num: %d\\\n\", 42)        // num: 42\\\nfmt.Printf(\"%s is %d\\\n\", \"age\", 30) // age is 30\\\nfmt.Printf(\"%.2f\\\n\", 3.14159)    // 3.14\\\nfmt.Printf(\"%v %T\\\n\", 42, 42)    // 42 int",
	UseCase:        "Use fmt.Printf when you need precise control over output formatting, building custom log messages, or creating formatted reports. Essential for debugging with detailed output. Use %v as a safe default for any type.",
	Prerequisites:  nil,
	RelatedTopics:  []string{"fmt-println"},
	DocsURL:        "https://pkg.go.dev/fmt",
}

func init() {
	Register(Concept065)
}
