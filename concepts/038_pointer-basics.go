package concepts

// 038. Pointer Basics
var Concept038 = Concept{
	Number:      38,
	ID:          "pointer-basics",
	Category:    "Pointers & Methods",
	Name:        "38. Pointer Basics",
	Description: "Get address with &",
	Instruction: "Create a variable x with value 42, create a pointer p that points to x, then print the dereferenced value of the pointer",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	x := 42
	p := &x
	fmt.Println(*p)
}`,
	ExpectedOutput: "42",
	Difficulty:     "beginner",
	Explanation:    "Pointers store memory addresses of values. The & operator gets the address of a variable (creates a pointer). The * operator dereferences a pointer (accesses the value). Go pointers are safer than C pointers - no pointer arithmetic.",
	Example:        "x := 42\np := &x       // p is *int, points to x\nfmt.Println(p)  // memory address\nfmt.Println(*p) // 42 (dereference)\n\nvar q *int      // nil pointer\nif q == nil {\n    fmt.Println(\"nil\")\n}",
	UseCase:        "Use pointers to avoid copying large structs, enable functions to modify arguments, share data between functions, and implement data structures like linked lists. Pointers are essential for efficient memory usage.",
	Prerequisites:  []string{"var-declaration"},
	RelatedTopics:  []string{"dereference", "pointer-struct", "value-vs-reference"},
	DocsURL:        "https://go.dev/tour/moretypes/1",
}

func init() {
	Register(Concept038)
}
