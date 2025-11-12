package concepts

// 029. Multiple Return Values
var Concept029 = Concept{
	Number:      29,
	ID:          "multiple-return",
	Category:    "Functions & Closures",
	Name:        "29. Multiple Return Values",
	Description: "Return multiple values from function",
	Instruction: "Implement a function named swap that takes two integers and returns them in reverse order",
	Boilerplate: `package main

import "fmt"

func swap(x, y int) (int, int) {
	// Your code here
}

func main() {
	a, _ := swap(1, 2)
	fmt.Println(a)
	c, _ := swap(10, 20)
	fmt.Println(c)
}`,
	Answer: `package main

import "fmt"

func swap(x, y int) (int, int) {
	return y, x
}

func main() {
	a, _ := swap(1, 2)
	fmt.Println(a)
	c, _ := swap(10, 20)
	fmt.Println(c)
}`,
	ExpectedOutput: "2\n20",
	TestCases: []TestCase{
		{Input: "swap(1, 2)", Expected: "2, 1"},
		{Input: "swap(10, 20)", Expected: "20, 10"},
	},
	Difficulty:     "beginner",
	Explanation:    "Go functions can return multiple values. This is commonly used to return a result and an error, or multiple related values. The return types are listed in parentheses. Callers receive all values using multiple assignment.",
	Example:        "func divide(a, b float64) (float64, error) {\n    if b == 0 {\n        return 0, errors.New(\"division by zero\")\n    }\n    return a / b, nil\n}\nresult, err := divide(10, 2)\nif err != nil {\n    // handle error\n}",
	UseCase:        "Use multiple returns primarily for error handling (value, error pattern). Also useful for returning related values like (quotient, remainder) or (min, max). This pattern makes error handling explicit and avoids exceptions.",
	Prerequisites:  []string{"function-params"},
	RelatedTopics:  []string{"error-check", "named-return", "multiple-assignment"},
	DocsURL:        "https://go.dev/tour/basics/6",
}

func init() {
	Register(Concept029)
}
