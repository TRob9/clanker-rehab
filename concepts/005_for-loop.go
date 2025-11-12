package concepts

// 005. For Loop
var Concept005 = Concept{
	Number:      5,
	ID:          "for-loop",
	Category:    "Core Syntax",
	Name:        "5. For Loop",
	Description: "Basic for loop with init, condition, post",
	Instruction: "Print numbers 1 to 3, one per line, using a for loop",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
}`,
	ExpectedOutput: "1\n2\n3",
	Difficulty:     "beginner",
	Explanation:    "Go has only one looping construct: the for loop. The classic for loop has three components: init (executed before first iteration), condition (evaluated before each iteration), and post (executed after each iteration). All three are optional.",
	Example:        "for i := 0; i < 5; i++ {\n    fmt.Println(i)\n}\n// Also works:\nfor i < 10 {  // while-style\n    i++\n}\nfor {  // infinite loop\n    break\n}",
	UseCase:        "Use for loops for iteration, repeated actions, and traversing collections. The classic three-part for is ideal when you know the number of iterations. Omit parts as needed for while-style or infinite loops.",
	Prerequisites:  []string{"var-declaration", "short-declaration"},
	RelatedTopics:  []string{"for-while", "range-slice", "if-else"},
	DocsURL:        "https://go.dev/tour/flowcontrol/1",
}

func init() {
	Register(Concept005)
}
