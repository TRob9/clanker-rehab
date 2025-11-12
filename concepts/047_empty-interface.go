package concepts

// 047. Empty Interface
var Concept047 = Concept{
	Number:      47,
	ID:          "empty-interface",
	Category:    "Interfaces",
	Name:        "47. Empty Interface",
	Description: "interface{} can hold any type",
	Instruction: "Declare a variable of type empty interface with value 42, then print it",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	var x interface{} = 42
	fmt.Println(x)
}`,
	ExpectedOutput: "42",
	Difficulty:     "beginner",
	Explanation:    "The empty interface interface{} has no methods, so every type satisfies it. It can hold values of any type. In Go 1.18+, 'any' is an alias for interface{}. Use for generic containers or when type is truly unknown.",
	Example:        "var x interface{}\nx = 42        // int\nx = \"hello\"   // string\nx = []int{1,2} // slice\n\n// Function accepting any type:\nfunc Print(v interface{}) {\n    fmt.Println(v)\n}",
	UseCase:        "Use empty interface sparingly - for truly generic code, JSON unmarshaling, fmt.Println, or when integrating with reflection. Prefer concrete types or generics (Go 1.18+) when possible for type safety.",
	Prerequisites:  []string{"interface-basic"},
	RelatedTopics:  []string{"type-assertion", "type-switch", "interface-basic"},
	DocsURL:        "https://go.dev/tour/methods/14",
}

func init() {
	Register(Concept047)
}
