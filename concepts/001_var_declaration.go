package concepts

// 001. Variable Declaration (var)
var Concept001 = Concept{
	Number:      1,
	ID:          "var-declaration",
	Category:    "Core Syntax",
	Name:        "1. Variable Declaration (var)",
	Description: "Declare variables using var keyword with explicit type",
	Instruction: "Declare an integer variable named x with the value 42, then print it",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	var x int = 42
	fmt.Println(x)
}`,
	ExpectedOutput: "42",
	Difficulty:     "beginner",
	Explanation:    "The 'var' keyword is used to declare variables with an explicit type. This is Go's most verbose but clearest way to declare variables. The syntax is: var name type = value. You can also declare without initializing: var name type (it gets the zero value).",
	Example:        "var age int = 25\nvar name string = \"Alice\"\nvar isActive bool = true\nvar count int // initialized to 0",
	UseCase:        "Use 'var' when you want to be explicit about types, when declaring package-level variables, or when you need to declare a variable without immediately initializing it. It's particularly useful for readability in larger codebases where type clarity matters.",
	Prerequisites:  nil,
	RelatedTopics:  []string{"short-declaration", "constants", "type-conversion"},
	DocsURL:        "https://go.dev/tour/basics/8",
}

func init() {
	Register(Concept001)
}
