package concepts

// 036. Defer Execution Order
var Concept036 = Concept{
	Number:      36,
	ID:          "defer-order",
	Category:    "Functions & Closures",
	Name:        "36. Defer Execution Order",
	Description: "Multiple defers execute in LIFO",
	Instruction: "Use defer to print 1, then defer to print 2, then print 3 directly, which should output 3, 2, 1 on separate lines",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}`,
	ExpectedOutput: "3\n2\n1",
	Difficulty:     "beginner",
	Explanation:    "When multiple defer statements exist, they execute in Last-In-First-Out (LIFO) order - like a stack. The most recently deferred function runs first. This ensures cleanup happens in reverse order of initialization.",
	Example:        "func process() {\n    fmt.Println(\"start\")\n    defer fmt.Println(\"cleanup 1\")\n    defer fmt.Println(\"cleanup 2\")\n    defer fmt.Println(\"cleanup 3\")\n    fmt.Println(\"work\")\n}\n// Prints: start, work, cleanup 3, cleanup 2, cleanup 1",
	UseCase:        "LIFO order is crucial for proper cleanup. If you acquire resources A then B, you should release B then A. Defer's automatic ordering handles this correctly: defer A.Release(); defer B.Release() will release B first, then A.",
	Prerequisites:  []string{"defer"},
	RelatedTopics:  []string{"defer", "panic-recover"},
	DocsURL:        "https://go.dev/blog/defer-panic-and-recover",
}

func init() {
	Register(Concept036)
}
