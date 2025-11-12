package concepts

// 012. Variable Shadowing
var Concept012 = Concept{
	Number:      12,
	ID:          "shadowing",
	Category:    "Core Syntax",
	Name:        "12. Variable Shadowing",
	Description: "Inner scope shadows outer variable",
	Instruction: "Declare a variable x with value 1, then in a new code block redeclare x with value 2 and print it, then print the outer x, which should print two then one",
	Boilerplate: `package main

import "fmt"

func main() {
	x := 1
	{
		x := 2
		fmt.Println(x)
	}
	fmt.Println(x)
}`,
	Answer: `package main

import "fmt"

func main() {
	x := 1
	{
		x := 2
		fmt.Println(x)
	}
	fmt.Println(x)
}`,
	ExpectedOutput: "2\n1",
	Difficulty:     "beginner",
	Explanation:    "Variable shadowing occurs when an inner scope declares a new variable with the same name as an outer scope variable. The inner declaration temporarily hides the outer one. Each scope sees its own version.",
	Example:        "x := \"outer\"\nif true {\n    x := \"inner\"  // shadows outer x\n    fmt.Println(x) // \"inner\"\n}\nfmt.Println(x) // \"outer\"\n// Common with := in if:\nif x, err := f(); err == nil {\n    // x here shadows outer x\n}",
	UseCase:        "Shadowing can be useful for limiting scope, but it's often a source of bugs. Be careful when using := in inner scopes - you might accidentally shadow instead of reassign. Tools like go vet can catch some shadowing issues.",
	Prerequisites:  []string{"short-declaration", "if-else"},
	RelatedTopics:  []string{"short-declaration", "if-else"},
	DocsURL:        "https://go.dev/doc/effective_go#redeclaration",
}

func init() {
	Register(Concept012)
}
