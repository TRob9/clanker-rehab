package concepts

// 088. Type Definition
var Concept088 = Concept{
	Number:      88,
	ID:          "type-definition",
	Category:    "Miscellaneous",
	Name:        "88. Type Definition",
	Description: "Define new type (not alias)",
	Instruction: "Create a new type MyInt based on int without using equals, declare a variable x of type MyInt with value 7, then print x",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	type MyInt int
	var x MyInt = 7
	fmt.Println(x)
}`,
	ExpectedOutput: "7",
	Difficulty:     "beginner",
	Explanation:    "Type definitions create new, distinct types based on existing types (without =). The new type has the same underlying representation but is not interchangeable with the original - explicit conversion is required. You can add methods to defined types.",
	Example:        "type Celsius float64\ntype Fahrenheit float64\n\nvar c Celsius = 100\n// var f Fahrenheit = c  // ERROR: different types\nvar f Fahrenheit = Fahrenheit(c)  // OK with conversion\n\n// Can add methods:\nfunc (c Celsius) ToF() Fahrenheit {\n    return Fahrenheit(c*9/5 + 32)\n}",
	UseCase:        "Use type definitions to create domain-specific types, add type safety, prevent mixing incompatible values (meters vs feet), or attach methods to types. Creates stronger type boundaries than aliases.",
	Prerequisites:  nil,
	RelatedTopics:  []string{"type-alias", "method-value-receiver"},
	DocsURL:        "https://go.dev/ref/spec#Type_declarations",
}

func init() {
	Register(Concept088)
}
