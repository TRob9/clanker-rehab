package concepts

// 094. Basic Generics (Go 1.18+)
var Concept094 = Concept{
	Number:      94,
	ID:          "generics-basic",
	Category:    "Miscellaneous",
	Name:        "94. Basic Generics (Go 1.18+)",
	Description: "Write a generic function that works with any type",
	Instruction: "Write a generic function named First with type parameter T constrained to any type that takes a slice of T and returns T by returning the first element. Call it with an integer slice containing 5 and 10, then print the result.",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	func First[T any](s []T) T { return s[0] }
	fmt.Println(First([]int{5, 10}))
}`,
	ExpectedOutput: "5",
	Difficulty:     "advanced",
	Explanation:    "Generics (type parameters) allow writing functions and types that work with any type. The syntax [T any] declares a type parameter T that can be any type. This enables type-safe reusable code without interface{} and type assertions.",
	Example:        "// Generic function:\nfunc First[T any](s []T) T {\n    return s[0]\n}\n\nFirst[int]([]int{1, 2, 3})       // 1\nFirst[string]([]string{\"a\", \"b\"}) // \"a\"\n// Type inference:\nFirst([]float64{1.5, 2.5})  // 1.5 (T inferred)\n\n// Generic type:\ntype Stack[T any] struct {\n    items []T\n}",
	UseCase:        "Use generics for type-safe data structures (stacks, queues, trees), algorithms that work on multiple types, or utility functions. Generics eliminate code duplication and provide compile-time type safety. Better than interface{} for performance and type safety.",
	Prerequisites:  []string{"interface-basic", "slice-basics"},
	RelatedTopics:  []string{"interface-type-constraint", "empty-interface"},
	DocsURL:        "https://go.dev/doc/tutorial/generics",
}

func init() {
	Register(Concept094)
}
