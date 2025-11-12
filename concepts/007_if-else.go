package concepts

// 007. If/Else
var Concept007 = Concept{
	Number:      7,
	ID:          "if-else",
	Category:    "Core Syntax",
	Name:        "7. If/Else",
	Description: "Conditional branching",
	Instruction: "If 10 is greater than 5, print "yes", otherwise print "no"",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	if 10 > 5 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}`,
	ExpectedOutput: "yes",
	Difficulty:     "beginner",
	Explanation:    "The if statement executes code when a condition is true. Optionally, add else or else if for alternative paths. Go's if can include a short statement before the condition, limiting variable scope to the if block.",
	Example:        "if x > 10 {\n    fmt.Println(\"big\")\n} else if x > 5 {\n    fmt.Println(\"medium\")\n} else {\n    fmt.Println(\"small\")\n}\n// With init:\nif v := getValue(); v > 0 {\n    fmt.Println(v)\n}",
	UseCase:        "Use if/else for conditional logic and decision-making. The init statement pattern is useful for limiting variable scope and checking error values immediately after function calls.",
	Prerequisites:  nil,
	RelatedTopics:  []string{"switch", "for-loop"},
	DocsURL:        "https://go.dev/tour/flowcontrol/5",
}

func init() {
	Register(Concept007)
}
