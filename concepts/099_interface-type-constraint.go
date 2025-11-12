package concepts

// 099. Interface as Type Constraint
var Concept099 = Concept{
	Number:      99,
	ID:          "interface-type-constraint",
	Category:    "Interfaces",
	Name:        "99. Interface as Type Constraint",
	Description: "Use interface to constrain generic type parameters",
	Instruction: "Define a Number interface with approximation operator for int and float64 using union syntax. Create a generic Double function that takes a Number and returns it multiplied by 2. Call it with 3 and print the result",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	type Number interface { ~int | ~float64 }
	func Double[T Number](n T) T { return n * 2 }
	fmt.Println(Double(3))
}`,
	ExpectedOutput: "6",
	Difficulty:     "advanced",
	Explanation:    "Go 1.18+ allows interfaces to constrain generic type parameters using union types (|) and approximation (~). This creates type constraints for generics. The ~ allows any type with the same underlying type.",
	Example:        "type Number interface {\n    ~int | ~float64  // int, float64, or custom types based on them\n}\n\nfunc Add[T Number](a, b T) T {\n    return a + b\n}\n\ntype MyInt int\nAdd(1, 2)              // int\nAdd(1.5, 2.5)          // float64\nAdd(MyInt(1), MyInt(2)) // MyInt",
	UseCase:        "Use type constraints for generic functions and types when you need operations on multiple types. Common for math functions, containers, and algorithms. Much better than interface{} + reflection for type-safe generic code.",
	Prerequisites:  []string{"interface-basic", "generics-basic"},
	RelatedTopics:  []string{"generics-basic", "interface-basic"},
	DocsURL:        "https://go.dev/doc/tutorial/generics",
}

func init() {
	Register(Concept099)
}
