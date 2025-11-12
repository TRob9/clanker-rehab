package concepts

// 028. Function with Parameters
var Concept028 = Concept{
	Number:      28,
	ID:          "function-params",
	Category:    "Functions & Closures",
	Name:        "28. Function with Parameters",
	Description: "Define function that accepts args",
	Instruction: "Implement a function named add that takes two integers and returns their sum",
	Boilerplate: `package main

import "fmt"

func add(a, b int) int {
	// Your code here
}

func main() {
	fmt.Println(add(2, 3))
	fmt.Println(add(10, 20))
	fmt.Println(add(-5, 5))
}`,
	Answer: `package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(2, 3))
	fmt.Println(add(10, 20))
	fmt.Println(add(-5, 5))
}`,
	ExpectedOutput: "5\n30\n0",
	TestCases: []TestCase{
		{Input: "add(2, 3)", Expected: "5"},
		{Input: "add(10, 20)", Expected: "30"},
		{Input: "add(-5, 5)", Expected: "0"},
	},
	Difficulty:     "beginner",
	Explanation:    "Functions can accept parameters to receive input values. Parameters are specified with name and type. Consecutive parameters of the same type can share the type declaration (a, b int). Functions use these parameters in their body.",
	Example:        "func add(a, b int) int {\n    return a + b\n}\nfunc greet(name string, age int) string {\n    return fmt.Sprintf(\"%s is %d\", name, age)\n}\nresult := add(2, 3)  // 5\nmsg := greet(\"Alice\", 30)",
	UseCase:        "Use function parameters to make functions reusable with different inputs. Parameters enable abstraction and code reuse. They're essential for building modular, testable code.",
	Prerequisites:  nil,
	RelatedTopics:  []string{"multiple-return", "variadic-function", "function-value"},
	DocsURL:        "https://go.dev/tour/basics/4",
}

func init() {
	Register(Concept028)
}
