package concepts

// 024. Struct Literal
var Concept024 = Concept{
	Number:      24,
	ID:          "struct-literal",
	Category:    "Data Structures",
	Name:        "24. Struct Literal",
	Description: "Initialize struct fields",
	Instruction: "Define a struct type called Point with X and Y fields, both integers, create an instance using positional initialization with values 1 and 2, then print the Y field",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	type Point struct { X, Y int }
	p := Point{1, 2}
	fmt.Println(p.Y)
}`,
	ExpectedOutput: "2",
	Difficulty:     "beginner",
	Explanation:    "Struct literals create and initialize struct values. Use field:value syntax for clarity or positional values. Unspecified fields get zero values. You can create pointers to structs using &StructLiteral{}.",
	Example:        "type Point struct { X, Y int }\np1 := Point{X: 1, Y: 2}  // named fields\np2 := Point{1, 2}         // positional\np3 := Point{X: 5}         // Y is 0\nptr := &Point{3, 4}       // pointer to struct",
	UseCase:        "Use struct literals for concise initialization, creating temporary values, or inline struct creation. Named field syntax is clearer and safer (field order independent). Use &StructLiteral{} when you need a pointer.",
	Prerequisites:  []string{"struct-declaration"},
	RelatedTopics:  []string{"struct-declaration", "pointer-struct"},
	DocsURL:        "https://go.dev/tour/moretypes/5",
}

func init() {
	Register(Concept024)
}
