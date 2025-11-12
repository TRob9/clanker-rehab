package concepts

// 086. Constant Block
var Concept086 = Concept{
	Number:      86,
	ID:          "const-block",
	Category:    "Miscellaneous",
	Name:        "86. Constant Block",
	Description: "Group constants in const() block",
	Instruction: "Create a constant block with constant A set to 1 and constant B set to 2, then print B",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	const (
		A = 1
		B = 2
	)
	fmt.Println(B)
}`,
	ExpectedOutput: "2",
	Difficulty:     "beginner",
	Explanation:    "Constant blocks group related constants together using const ( ... ) syntax. Inside blocks, uninitialized constants repeat the previous value and type. Particularly useful with iota for creating enumerations.",
	Example:        "const (\n    StatusOK = 200\n    StatusCreated = 201\n    StatusBadRequest = 400\n)\n\n// With iota:\nconst (\n    Read = 1 << iota  // 1\n    Write             // 2 (repeated expression)\n    Execute           // 4\n)",
	UseCase:        "Use const blocks to group related constants for organization and readability. Common for status codes, configuration values, bit flags, or enumerations. Makes code more maintainable by keeping related constants together.",
	Prerequisites:  []string{"constants"},
	RelatedTopics:  []string{"constants", "iota"},
	DocsURL:        "https://go.dev/ref/spec#Constant_declarations",
}

func init() {
	Register(Concept086)
}
