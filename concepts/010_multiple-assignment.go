package concepts

// 010. Multiple Assignment
var Concept010 = Concept{
	Number:      10,
	ID:          "multiple-assignment",
	Category:    "Core Syntax",
	Name:        "10. Multiple Assignment",
	Description: "Assign multiple variables in one line",
	Instruction: "Use multiple assignment to assign 1 to variable a and 2 to variable b in one line, then print the sum of the two variables",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Println(a + b)
}`,
	ExpectedOutput: "3",
	Difficulty:     "beginner",
	Explanation:    "Go allows assigning multiple variables simultaneously. This is useful for swapping values, unpacking multiple return values, and concise initialization of related variables.",
	Example:        "a, b := 1, 2\n// Swap without temp:\na, b = b, a\n// From function:\nval, err := someFunc()\n// Multiple returns:\nx, y, z := 1, 2, 3",
	UseCase:        "Use multiple assignment for swapping variables without a temp, receiving multiple return values (especially value/error pairs), and initializing related variables together.",
	Prerequisites:  []string{"short-declaration"},
	RelatedTopics:  []string{"multiple-return", "blank-identifier"},
	DocsURL:        "https://go.dev/tour/basics/15",
}

func init() {
	Register(Concept010)
}
