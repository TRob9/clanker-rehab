package concepts

// 087. Type Alias
var Concept087 = Concept{
	Number:      87,
	ID:          "type-alias",
	Category:    "Miscellaneous",
	Name:        "87. Type Alias",
	Description: "Create type alias with =",
	Instruction: "Create a type alias MyInt for int using the equals syntax, declare a variable x of type MyInt with value 5, then print x",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	type MyInt = int
	var x MyInt = 5
	fmt.Println(x)
}`,
	ExpectedOutput: "5",
	Difficulty:     "beginner",
	Explanation:    "Type aliases create alternate names for existing types using the = syntax. The alias and original type are completely interchangeable - they're the same type. Unlike type definitions, no conversion is needed.",
	Example:        "type MyInt = int  // alias\n\nvar x MyInt = 5\nvar y int = x  // No conversion needed\n\n// Common use:\ntype StringSet = map[string]struct{}\n\n// Byte and rune are aliases:\ntype byte = uint8\ntype rune = int32",
	UseCase:        "Use type aliases for gradual code refactoring, providing clearer names without changing behavior, or when migrating between types. Unlike type definitions, aliases don't create new types - they're just different names for the same type.",
	Prerequisites:  nil,
	RelatedTopics:  []string{"type-definition"},
	DocsURL:        "https://go.dev/ref/spec#Type_declarations",
}

func init() {
	Register(Concept087)
}
