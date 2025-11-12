package concepts

// 034. Recursive Function
var Concept034 = Concept{
	Number:      34,
	ID:          "recursion",
	Category:    "Functions & Closures",
	Name:        "34. Recursive Function",
	Description: "Function that calls itself",
	Instruction: "Implement a function named factorial that calculates the factorial of n recursively",
	Boilerplate: `package main

import "fmt"

func factorial(n int) int {
	// Your code here
}

func main() {
	fmt.Println(factorial(3))
	fmt.Println(factorial(5))
	fmt.Println(factorial(0))
}`,
	Answer: `package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(3))
	fmt.Println(factorial(5))
	fmt.Println(factorial(0))
}`,
	ExpectedOutput: "6\n120\n1",
	TestCases: []TestCase{
		{Input: "factorial(3)", Expected: "6"},
		{Input: "factorial(5)", Expected: "120"},
		{Input: "factorial(0)", Expected: "1"},
	},
	Difficulty:     "beginner",
	Explanation:    "Recursion is when a function calls itself. Recursive functions need a base case (termination condition) to prevent infinite recursion. Each recursive call creates a new stack frame until the base case is reached.",
	Example:        "func factorial(n int) int {\n    if n <= 1 {  // base case\n        return 1\n    }\n    return n * factorial(n-1)  // recursive case\n}\n// Tree traversal:\nfunc walk(node *Node) {\n    if node == nil { return }\n    walk(node.Left)\n    process(node)\n    walk(node.Right)\n}",
	UseCase:        "Use recursion for naturally recursive problems: tree/graph traversal, divide-and-conquer algorithms, backtracking. For simple iteration, prefer loops. Watch for stack overflow with deep recursion - consider iterative solutions for large inputs.",
	Prerequisites:  []string{"function-params"},
	RelatedTopics:  []string{"function-params"},
	DocsURL:        "https://go.dev/tour/flowcontrol/9",
}

func init() {
	Register(Concept034)
}
