package concepts

// 049. Type Switch
var Concept049 = Concept{
	Number:      49,
	ID:          "type-switch",
	Category:    "Interfaces",
	Name:        "49. Type Switch",
	Description: "Switch on interface type",
	Instruction: "Declare an empty interface variable with value 3, then use a type switch to print "int"" if it is an integer",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	var i interface{} = 3
	switch i.(type) {
	case int:
		fmt.Println("int")
	}
}`,
	ExpectedOutput: "int",
	Difficulty:     "beginner",
	Explanation:    "Type switches allow switching on the type of an interface value. The special syntax switch v := i.(type) assigns the concrete value to v in each case. Each case can handle a different type.",
	Example:        "func describe(i interface{}) {\n    switch v := i.(type) {\n    case int:\n        fmt.Printf(\"int: %d\\\n\", v)\n    case string:\n        fmt.Printf(\"string: %s\\\n\", v)\n    default:\n        fmt.Printf(\"unknown: %T\\\n\", v)\n    }\n}",
	UseCase:        "Use type switches when you need to handle multiple possible types from an interface, implement polymorphic behavior, or work with empty interfaces. Cleaner than multiple type assertions.",
	Prerequisites:  []string{"type-assertion", "switch"},
	RelatedTopics:  []string{"type-assertion", "empty-interface", "switch"},
	DocsURL:        "https://go.dev/tour/methods/16",
}

func init() {
	Register(Concept049)
}
