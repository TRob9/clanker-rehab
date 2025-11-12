package concepts

// 030. Named Return Values
var Concept030 = Concept{
	Number:      30,
	ID:          "named-return",
	Category:    "Functions & Closures",
	Name:        "30. Named Return Values",
	Description: "Use named returns",
	Instruction: "Implement a function that returns a named integer variable with value 7",
	Boilerplate: `package main

import "fmt"

func f() (x int) {
	// Your code here
}

func main() {
	fmt.Println(f())
}`,
	Answer: `package main

import "fmt"

func f() (x int) {
	x = 7
	return
}

func main() {
	fmt.Println(f())
}`,
	ExpectedOutput: "7",
	TestCases: []TestCase{
		{Input: "f()", Expected: "7"},
	},
	Difficulty:     "beginner",
	Explanation:    "Named return values are pre-declared variables in the function signature. They're initialized to zero values and can be modified in the function body. A bare 'return' statement returns the current values of named returns.",
	Example:        "func split(sum int) (x, y int) {\n    x = sum * 4 / 9\n    y = sum - x\n    return  // returns x and y\n}\n// Also useful for clarity:\nfunc ReadFile() (data []byte, err error) {\n    data, err = os.ReadFile(\"file.txt\")\n    return\n}",
	UseCase:        "Use named returns to document what the function returns, especially in longer functions. They can improve readability and are useful with defer for cleanup. However, use sparingly - they can make code less clear if overused.",
	Prerequisites:  []string{"multiple-return"},
	RelatedTopics:  []string{"multiple-return", "defer"},
	DocsURL:        "https://go.dev/tour/basics/7",
}

func init() {
	Register(Concept030)
}
