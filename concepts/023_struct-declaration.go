package concepts

// 023. Struct Declaration
var Concept023 = Concept{
	Number:      23,
	ID:          "struct-declaration",
	Category:    "Data Structures",
	Name:        "23. Struct Declaration",
	Description: "Define and use struct",
	Instruction: "Define a struct type called Person with a Name field of type string, create an instance with the Name field set to Alice, then print the Name field",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	type Person struct {
		Name string
	}
	p := Person{Name: "Alice"}
	fmt.Println(p.Name)
}`,
	ExpectedOutput: "Alice",
	Difficulty:     "beginner",
	Explanation:    "Structs are typed collections of fields. They group related data together to form records. Each field has a name and type. Structs are value types - they're copied when assigned or passed to functions.",
	Example:        "type Person struct {\n    Name string\n    Age  int\n}\np := Person{Name: \"Alice\", Age: 30}\np2 := Person{\"Bob\", 25}  // positional",
	UseCase:        "Use structs to model entities, create records, group related data, or define custom types. Structs are the foundation of object-oriented patterns in Go, combined with methods.",
	Prerequisites:  []string{"var-declaration"},
	RelatedTopics:  []string{"struct-literal", "method-value-receiver", "method-pointer-receiver"},
	DocsURL:        "https://go.dev/tour/moretypes/2",
}

func init() {
	Register(Concept023)
}
