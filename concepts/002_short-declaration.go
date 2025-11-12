package concepts

// 002. Short Variable Declaration
var Concept002 = Concept{
	Number:      2,
	ID:          "short-declaration",
	Category:    "Core Syntax",
	Name:        "2. Short Variable Declaration",
	Description: "Use := for short variable declaration (type inferred)",
	Instruction: "Declare a variable y with the value hello using short declaration syntax and print it",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	y := "hello"
	fmt.Println(y)
}`,
	ExpectedOutput: "hello",
	Difficulty:     "beginner",
	Explanation:    "The := operator is Go's shorthand for declaring and initializing variables. The compiler automatically infers the type from the value. This is the most common way to declare variables inside functions. You cannot use := outside of functions.",
	Example:        "name := \"Bob\"        // inferred as string\nage := 30            // inferred as int\nisActive := true     // inferred as bool\nvalues := []int{1,2} // inferred as []int slice",
	UseCase:        "Use := for almost all variable declarations inside functions. It's concise, idiomatic Go, and reduces boilerplate. However, you must use 'var' for package-level variables and when you need to declare without initializing.",
	Prerequisites:  []string{"var-declaration"},
	RelatedTopics:  []string{"var-declaration", "type-conversion", "multiple-assignment"},
	DocsURL:        "https://go.dev/tour/basics/10",
}

func init() {
	Register(Concept002)
}
