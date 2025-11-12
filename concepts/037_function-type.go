package concepts

// 037. Function Type
var Concept037 = Concept{
	Number:      37,
	ID:          "function-type",
	Category:    "Functions & Closures",
	Name:        "37. Function Type",
	Description: "Define function type",
	Instruction: "Define a function type called Op that takes two integers and returns an integer, then create a variable of that type that adds the two numbers, then print the result of calling it with 4 and 5",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	type Op func(int, int) int
	add := Op(func(a, b int) int { return a + b })
	fmt.Println(add(4, 5))
}`,
	ExpectedOutput: "9",
	Difficulty:     "intermediate",
	Explanation:    "You can define custom types for function signatures using 'type'. This makes function signatures reusable, improves readability, and allows attaching methods to function types. Function types can be used anywhere a type is expected.",
	Example:        "type BinaryOp func(int, int) int\n\nfunc apply(op BinaryOp, a, b int) int {\n    return op(a, b)\n}\n\nvar add BinaryOp = func(a, b int) int { return a + b }\nvar mul BinaryOp = func(a, b int) int { return a * b }\n\napply(add, 2, 3)  // 5\napply(mul, 2, 3)  // 6",
	UseCase:        "Use function types for callbacks, strategy patterns, middleware chains, and when the same function signature is used repeatedly. Common in HTTP handlers (http.HandlerFunc), middleware, and plugin systems.",
	Prerequisites:  []string{"function-value"},
	RelatedTopics:  []string{"function-value", "closure"},
	DocsURL:        "https://go.dev/ref/spec#Function_types",
}

func init() {
	Register(Concept037)
}
