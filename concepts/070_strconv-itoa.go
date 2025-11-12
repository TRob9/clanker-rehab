package concepts

// 070. strconv.Itoa
var Concept070 = Concept{
	Number:      70,
	ID:          "strconv-itoa",
	Category:    "Standard Library",
	Name:        "70. strconv.Itoa",
	Description: "Convert int to string",
	Instruction: "Use the Itoa function from the strconv package to convert the integer 456 to a string, then print the result",
	Boilerplate: `package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Your code here
}`,
	Answer: `package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := strconv.Itoa(456)
	fmt.Println(s)
}`,
	ExpectedOutput: "456",
	Difficulty:     "beginner",
	Explanation:    "strconv.Itoa (integer to ASCII) converts an int to its decimal string representation. Unlike Atoi, it doesn't return an error - int to string conversion always succeeds.",
	Example:        "strconv.Itoa(123)   // \"123\"\nstrconv.Itoa(-456)  // \"-456\"\nstrconv.Itoa(0)     // \"0\"\n// Building strings:\nmsg := \"Count: \" + strconv.Itoa(42)  // \"Count: 42\"",
	UseCase:        "Use strconv.Itoa to convert integers to strings for display, logging, building URLs/queries, or string concatenation. For formatted output, fmt.Sprintf may be clearer. For other bases (hex, binary), use strconv.FormatInt.",
	Prerequisites:  nil,
	RelatedTopics:  []string{"strconv-atoi"},
	DocsURL:        "https://pkg.go.dev/strconv#Itoa",
}

func init() {
	Register(Concept070)
}
