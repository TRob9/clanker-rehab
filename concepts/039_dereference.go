package concepts

// 039. Dereferencing Pointers
var Concept039 = Concept{
	Number:      39,
	ID:          "dereference",
	Category:    "Pointers & Methods",
	Name:        "39. Dereferencing Pointers",
	Description: "Access value with *",
	Instruction: "Create a variable x with value 10, create a pointer p to x, use the dereference operator to set the value at p to 20, then print x",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	x := 10
	p := &x
	*p = 20
	fmt.Println(x)
}`,
	ExpectedOutput: "20",
	Difficulty:     "beginner",
	Explanation:    "Dereferencing a pointer means accessing the value it points to using the * operator. You can both read and write through a dereferenced pointer. Changes through the pointer affect the original variable.",
	Example:        "x := 10\np := &x\nfmt.Println(*p)  // 10 (read)\n*p = 20          // write\nfmt.Println(x)   // 20 (x changed)\n\n// Common pattern:\nfunc double(p *int) {\n    *p = *p * 2  // modifies original\n}",
	UseCase:        "Use dereferencing to modify values through pointers, implement functions that update their arguments, and work with optional values (nil pointers). Essential for building mutable data structures.",
	Prerequisites:  []string{"pointer-basics"},
	RelatedTopics:  []string{"pointer-basics", "value-vs-reference", "pointer-mutation"},
	DocsURL:        "https://go.dev/tour/moretypes/1",
}

func init() {
	Register(Concept039)
}
