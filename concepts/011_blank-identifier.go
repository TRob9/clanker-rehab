package concepts

// 011. Blank Identifier
var Concept011 = Concept{
	Number:      11,
	ID:          "blank-identifier",
	Category:    "Core Syntax",
	Name:        "11. Blank Identifier",
	Description: "Use _ to ignore values",
	Instruction: "Use the blank identifier to assign 10 to an ignored variable and 20 to variable b, then print variable b",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	_, b := 10, 20
	fmt.Println(b)
}`,
	ExpectedOutput: "20",
	Difficulty:     "beginner",
	Explanation:    "The blank identifier '_' is a write-only placeholder that discards values. It's used when Go requires a variable but you don't need the value. Unlike regular identifiers, _ can be assigned to repeatedly without declaring it.",
	Example:        "// Ignore error:\nvalue, _ := strconv.Atoi(\"123\")\n// Ignore index in range:\nfor _, v := range slice {\n    fmt.Println(v)\n}\n// Import for side effects:\nimport _ \"database/sql\"",
	UseCase:        "Use _ to ignore unwanted return values (especially errors when you're sure they won't occur), discard loop indices, or import packages only for their side effects (like driver registration).",
	Prerequisites:  []string{"multiple-assignment"},
	RelatedTopics:  []string{"multiple-return", "range-slice", "blank-import"},
	DocsURL:        "https://go.dev/doc/effective_go#blank",
}

func init() {
	Register(Concept011)
}
