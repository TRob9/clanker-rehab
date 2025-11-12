package concepts

// 052. Stringer Interface
var Concept052 = Concept{
	Number:      52,
	ID:          "stringer",
	Category:    "Interfaces",
	Name:        "52. Stringer Interface",
	Description: "Implement fmt.Stringer",
	Instruction: "Define a Person struct with a Name field. Implement the String method that returns the Name field. Create a Person with name Bob and print it",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

type Person struct { Name string }

func (p Person) String() string { return p.Name }

func main() {
	p := Person{Name: "Bob"}
	fmt.Println(p)
}`,
	ExpectedOutput: "Bob",
	Difficulty:     "beginner",
	Explanation:    "The fmt.Stringer interface has one method: String() string. Types implementing Stringer control how they're printed by fmt.Print, fmt.Println, and fmt.Sprintf. This is Go's equivalent of toString() in other languages.",
	Example:        "type Point struct { X, Y int }\n\nfunc (p Point) String() string {\n    return fmt.Sprintf(\"(%d, %d)\", p.X, p.Y)\n}\n\np := Point{3, 4}\nfmt.Println(p)  // \"(3, 4)\" instead of \"{3 4}\"",
	UseCase:        "Implement Stringer for custom types to provide readable string representations for logging, debugging, and display. Essential for types you'll print frequently. Makes fmt.Println output much more useful.",
	Prerequisites:  []string{"interface-basic", "method-value-receiver"},
	RelatedTopics:  []string{"interface-basic", "fmt-println"},
	DocsURL:        "https://pkg.go.dev/fmt#Stringer",
}

func init() {
	Register(Concept052)
}
