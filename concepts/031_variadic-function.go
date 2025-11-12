package concepts

// 031. Variadic Function
var Concept031 = Concept{
	Number:      31,
	ID:          "variadic-function",
	Category:    "Functions & Closures",
	Name:        "31. Variadic Function",
	Description: "Accept variable number of args",
	Instruction: "Implement a function named sum that accepts any number of integers using variadic parameters and returns their sum",
	Boilerplate: `package main

import "fmt"

func sum(nums ...int) int {
	// Your code here
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())
}`,
	Answer: `package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())
}`,
	ExpectedOutput: "6\n30\n0",
	TestCases: []TestCase{
		{Input: "sum(1, 2, 3)", Expected: "6"},
		{Input: "sum(10, 20)", Expected: "30"},
		{Input: "sum()", Expected: "0"},
	},
	Difficulty:     "beginner",
	Explanation:    "Variadic functions accept a variable number of arguments of the same type using ...T syntax. Inside the function, the variadic parameter is a slice. You can call with zero or more arguments, or expand a slice with ... operator.",
	Example:        "func sum(nums ...int) int {\n    total := 0\n    for _, n := range nums {\n        total += n\n    }\n    return total\n}\nsum(1, 2, 3)  // 6\nvalues := []int{1,2,3,4}\nsum(values...)  // expand slice",
	UseCase:        "Use variadic functions when the number of arguments is unknown or varies. Common in Printf-style functions, builders, and utility functions. The variadic parameter must be the last parameter.",
	Prerequisites:  []string{"function-params", "slice-basics"},
	RelatedTopics:  []string{"function-params", "range-slice"},
	DocsURL:        "https://go.dev/ref/spec#Passing_arguments_to_..._parameters",
}

func init() {
	Register(Concept031)
}
