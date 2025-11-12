package concepts

// 006. For as While
var Concept006 = Concept{
	Number:      6,
	ID:          "for-while",
	Category:    "Core Syntax",
	Name:        "6. For as While",
	Description: "Use for loop like a while loop",
	Instruction: "Start a counter at 1, loop while the counter is less than or equal to 2, print the counter each iteration, then increment the counter",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	i := 1
	for i <= 2 {
		fmt.Println(i)
		i++
	}
}`,
	ExpectedOutput: "1\n2",
	Difficulty:     "beginner",
	Explanation:    "Go doesn't have a while keyword. Instead, you can omit the init and post statements in a for loop, leaving only the condition. This creates a while-style loop that continues as long as the condition is true.",
	Example:        "i := 0\nfor i < 5 {  // Like while(i < 5)\n    fmt.Println(i)\n    i++\n}",
	UseCase:        "Use for-as-while when you don't need init/post statements, such as reading from a stream until EOF, waiting for a condition to become true, or when the loop variable is managed elsewhere.",
	Prerequisites:  []string{"for-loop"},
	RelatedTopics:  []string{"for-loop", "if-else"},
	DocsURL:        "https://go.dev/tour/flowcontrol/3",
}

func init() {
	Register(Concept006)
}
