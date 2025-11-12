package concepts

// 092. Type Assertion with ok
var Concept092 = Concept{
	Number:      92,
	ID:          "type-assertion-ok",
	Category:    "Miscellaneous",
	Name:        "92. Type Assertion with ok",
	Description: "Safely assert type with comma ok",
	Instruction: "Declare an empty interface with value 42, then use type assertion with the comma-ok idiom to safely check" if it is an integer. Print the ok value.",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	var i interface{} = 42
	v, ok := i.(int)
	_ = v
	fmt.Println(ok)
}`,
	ExpectedOutput: "true",
	Difficulty:     "beginner",
	Explanation:    "The comma-ok idiom for type assertions provides a safe way to check if an interface holds a specific type. If the assertion succeeds, ok is true and v holds the value. If it fails, ok is false and v is the zero value (no panic).",
	Example:        "var i interface{} = \"hello\"\n\n// Safe assertion:\nif s, ok := i.(string); ok {\n    fmt.Println(\"string:\", s)  // executes\n}\n\n// Failed assertion (no panic):\nif n, ok := i.(int); ok {\n    fmt.Println(\"int:\", n)\n} else {\n    fmt.Println(\"not an int\")  // executes\n}",
	UseCase:        "Always use the comma-ok form for type assertions unless you're 100% certain of the type. This prevents panics from failed assertions. Common in handling interface{} values, parsing JSON into specific types, or implementing type-specific behavior.",
	Prerequisites:  []string{"type-assertion", "empty-interface"},
	RelatedTopics:  []string{"type-assertion", "type-switch", "map-existence"},
	DocsURL:        "https://go.dev/tour/methods/15",
}

func init() {
	Register(Concept092)
}
