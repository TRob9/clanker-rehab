package concepts

// 066. strings.Contains
var Concept066 = Concept{
	Number:      66,
	ID:          "strings-contains",
	Category:    "Standard Library",
	Name:        "66. strings.Contains",
	Description: "Check if string contains substring",
	Instruction: "Use the Contains function from the strings package to check whether the word "hello" contains the substring ll, then print the result",
	Boilerplate: `package main

import (
	"fmt"
	"strings"
)

func main() {
	// Your code here
}`,
	Answer: `package main

import (
	"fmt"
	"strings"
)

func main() {
	result := strings.Contains("hello", "ll")
	fmt.Println(result)
}`,
	ExpectedOutput: "true",
	Difficulty:     "beginner",
	Explanation:    "strings.Contains checks if a string contains a substring, returning a boolean. It performs a simple substring search and is case-sensitive. Returns true if substr is empty.",
	Example:        "strings.Contains(\"hello\", \"ll\")    // true\nstrings.Contains(\"hello\", \"LL\")    // false (case-sensitive)\nstrings.Contains(\"hello\", \"\")      // true (empty substr)\nstrings.Contains(\"hello\", \"world\") // false",
	UseCase:        "Use strings.Contains for quick substring checks, input validation, filtering, or search functionality. For case-insensitive search, convert both strings to lowercase first. For finding position, use strings.Index instead.",
	Prerequisites:  nil,
	RelatedTopics:  []string{"strings-split"},
	DocsURL:        "https://pkg.go.dev/strings#Contains",
}

func init() {
	Register(Concept066)
}
