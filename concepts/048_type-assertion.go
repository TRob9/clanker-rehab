package concepts

// 048. Type Assertion
var Concept048 = Concept{
	Number:      48,
	ID:          "type-assertion",
	Category:    "Interfaces",
	Name:        "48. Type Assertion",
	Description: "Assert concrete type from interface",
	Instruction: "Declare an empty interface variable with value hello, use type assertion to extract it as a string variable, then print the string variable",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
}`,
	ExpectedOutput: "hello",
	Difficulty:     "beginner",
	Explanation:    "Type assertion extracts the concrete value from an interface. Syntax: value := interfaceVar.(Type). If the assertion fails, it panics. Use comma-ok idiom for safe assertions: value, ok := interfaceVar.(Type).",
	Example:        "var i interface{} = \"hello\"\ns := i.(string)  // \"hello\", panics if wrong type\n\n// Safe version:\nif s, ok := i.(string); ok {\n    fmt.Println(s)  // \"hello\"\n} else {\n    fmt.Println(\"not a string\")\n}",
	UseCase:        "Use type assertions to extract concrete types from interfaces, handle different types differently, or work with values from empty interfaces. Always use comma-ok form unless you're absolutely certain of the type.",
	Prerequisites:  []string{"empty-interface"},
	RelatedTopics:  []string{"empty-interface", "type-switch", "type-assertion-ok"},
	DocsURL:        "https://go.dev/tour/methods/15",
}

func init() {
	Register(Concept048)
}
