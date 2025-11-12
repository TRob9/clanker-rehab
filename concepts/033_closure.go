package concepts

// 033. Closures
var Concept033 = Concept{
	Number:      33,
	ID:          "closure",
	Category:    "Functions & Closures",
	Name:        "33. Closures",
	Description: "Function that captures outer variable",
	Instruction: "Create a variable x with value 10, then create a function that returns x plus 1, then call the function and print the result",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	x := 10
	f := func() int { return x + 1 }
	fmt.Println(f())
}`,
	ExpectedOutput: "11",
	Difficulty:     "beginner",
	Explanation:    "A closure is a function that references variables from outside its body. The function 'closes over' these variables, keeping them alive even after the outer function returns. Each closure maintains its own copy of captured variables.",
	Example:        "func counter() func() int {\n    count := 0\n    return func() int {\n        count++  // captures count\n        return count\n    }\n}\nc1 := counter()\nc1()  // 1\nc1()  // 2\nc2 := counter()  // independent\nc2()  // 1",
	UseCase:        "Use closures for encapsulation, maintaining state without global variables, creating function generators, callbacks that need context, and implementing iterators. Closures are powerful for creating clean, self-contained functions.",
	Prerequisites:  []string{"function-value"},
	RelatedTopics:  []string{"function-value", "function-type"},
	DocsURL:        "https://go.dev/tour/moretypes/25",
}

func init() {
	Register(Concept033)
}
