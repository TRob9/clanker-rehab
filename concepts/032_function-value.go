package concepts

// 032. Function as Value
var Concept032 = Concept{
	Number:      32,
	ID:          "function-value",
	Category:    "Functions & Closures",
	Name:        "32. Function as Value",
	Description: "Assign function to variable",
	Instruction: "Create a function variable called double that takes an integer and returns the integer multiplied by 2, then call it with 3 and print the result",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	double := func(x int) int { return x * 2 }
	fmt.Println(double(3))
}`,
	ExpectedOutput: "6",
	TestCases: []TestCase{
		{Input: "double(3)", Expected: "6"},
		{Input: "double(5)", Expected: "10"},
		{Input: "double(0)", Expected: "0"},
	},
	Difficulty:     "beginner",
	Explanation:    "Functions are first-class values in Go - they can be assigned to variables, passed as arguments, and returned from other functions. Anonymous functions (function literals) can be defined inline without a name.",
	Example:        "// Assign to variable:\ndouble := func(x int) int { return x * 2 }\nresult := double(3)  // 6\n// Pass as argument:\nfunc apply(f func(int) int, val int) int {\n    return f(val)\n}\napply(double, 5)  // 10",
	UseCase:        "Use function values for callbacks, strategy patterns, configurable behavior, and functional programming patterns. Common in sorting (custom comparators), HTTP handlers, and event handling.",
	Prerequisites:  []string{"function-params"},
	RelatedTopics:  []string{"closure", "function-type"},
	DocsURL:        "https://go.dev/tour/moretypes/24",
}

func init() {
	Register(Concept032)
}
