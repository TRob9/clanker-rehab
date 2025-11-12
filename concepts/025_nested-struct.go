package concepts

// 025. Nested Struct
var Concept025 = Concept{
	Number:      25,
	ID:          "nested-struct",
	Category:    "Data Structures",
	Name:        "25. Nested Struct",
	Description: "Struct with embedded struct",
	Instruction: "Define a type Address with field City of type string, define type Person with field A of type Address, create an instance p with the nested City field set to NYC, then print the nested City field through p",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	type Address struct { City string }
	type Person struct { A Address }
	p := Person{A: Address{City: "NYC"}}
	fmt.Println(p.A.City)
}`,
	ExpectedOutput: "NYC",
	Difficulty:     "beginner",
	Explanation:    "Structs can contain other structs as fields, creating nested or composite types. This allows modeling complex hierarchical data. Access nested fields using dot notation: outer.inner.field.",
	Example:        "type Address struct {\n    City, State string\n}\ntype Person struct {\n    Name string\n    Addr Address\n}\np := Person{\n    Name: \"Alice\",\n    Addr: Address{City: \"NYC\", State: \"NY\"},\n}\nfmt.Println(p.Addr.City)",
	UseCase:        "Use nested structs to model hierarchical data (person has address), compose complex types from simpler ones, or represent tree/graph structures. Nesting clarifies relationships and improves code organization.",
	Prerequisites:  []string{"struct-declaration"},
	RelatedTopics:  []string{"struct-declaration", "struct-embedding"},
	DocsURL:        "https://go.dev/tour/moretypes/2",
}

func init() {
	Register(Concept025)
}
