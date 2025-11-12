package concepts

// 003. Constants
var Concept003 = Concept{
	Number:      3,
	ID:          "constants",
	Category:    "Core Syntax",
	Name:        "3. Constants",
	Description: "Declare constants with const",
	Instruction: "Declare a constant Pi with the value 3.14 and print it",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	const Pi = 3.14
	fmt.Println(Pi)
}`,
	ExpectedOutput: "3.14",
	Difficulty:     "beginner",
	Explanation:    "Constants are immutable values declared with the 'const' keyword. Unlike variables, their values cannot be changed after declaration. Constants can be character, string, boolean, or numeric values. They're computed at compile time.",
	Example:        "const Pi = 3.14159\nconst AppName = \"MyApp\"\nconst MaxUsers = 100\nconst Enabled = true",
	UseCase:        "Use constants for values that never change throughout your program: configuration values, mathematical constants, fixed limits, or magic numbers. Constants make code more readable and prevent accidental modification of important values.",
	Prerequisites:  []string{"var-declaration"},
	RelatedTopics:  []string{"var-declaration", "iota", "const-block"},
	DocsURL:        "https://go.dev/tour/basics/15",
}

func init() {
	Register(Concept003)
}
