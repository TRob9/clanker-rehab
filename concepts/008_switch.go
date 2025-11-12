package concepts

// 008. Switch Statement
var Concept008 = Concept{
	Number:      8,
	ID:          "switch",
	Category:    "Core Syntax",
	Name:        "8. Switch Statement",
	Description: "Switch with cases",
	Instruction: "Create a switch statement with short declaration that assigns the value 2 to variable x, then print "two" when the case matches 2",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	switch x := 2; x {
	case 2:
		fmt.Println("two")
	}
}`,
	ExpectedOutput: "two",
	Difficulty:     "beginner",
	Explanation:    "Switch statements provide clean multi-way branching. Unlike C/Java, Go's switch breaks automatically (no fallthrough). Cases can be expressions, not just constants. You can also switch without a condition (like if-else chains).",
	Example:        "switch day {\ncase \"Mon\", \"Tue\":\n    fmt.Println(\"weekday\")\ncase \"Sat\", \"Sun\":\n    fmt.Println(\"weekend\")\ndefault:\n    fmt.Println(\"unknown\")\n}",
	UseCase:        "Use switch for multi-way branching when you have multiple conditions on the same value. Switch is cleaner than multiple if-else statements and automatically breaks, reducing bugs.",
	Prerequisites:  []string{"if-else"},
	RelatedTopics:  []string{"if-else", "type-switch"},
	DocsURL:        "https://go.dev/tour/flowcontrol/9",
}

func init() {
	Register(Concept008)
}
