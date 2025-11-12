package concepts

// 026. Anonymous Struct
var Concept026 = Concept{
	Number:      26,
	ID:          "anonymous-struct",
	Category:    "Data Structures",
	Name:        "26. Anonymous Struct",
	Description: "Struct without named type",
	Instruction: "Create an anonymous struct with field Val of type int, initialize it with Val set to 5, then print the Val field",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	x := struct{ Val int }{Val: 5}
	fmt.Println(x.Val)
}`,
	ExpectedOutput: "5",
	Difficulty:     "beginner",
	Explanation:    "Anonymous structs are struct types declared without a named type. They're useful for one-off data structures, temporary grouping, or JSON marshaling. The type is defined inline where used.",
	Example:        "person := struct {\n    Name string\n    Age  int\n}{Name: \"Alice\", Age: 30}\n// Common in tests:\ntests := []struct{\n    input int\n    want  int\n}{\n    {1, 2},\n    {2, 4},\n}",
	UseCase:        "Use anonymous structs for temporary data structures, table-driven tests, JSON unmarshaling into throwaway types, or when a full type definition would be overkill. They're perfect for one-off groupings.",
	Prerequisites:  []string{"struct-declaration"},
	RelatedTopics:  []string{"struct-literal", "struct-declaration"},
	DocsURL:        "https://go.dev/tour/moretypes/5",
}

func init() {
	Register(Concept026)
}
