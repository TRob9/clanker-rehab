package concepts

// 081. Panic and Recover
var Concept081 = Concept{
	Number:      81,
	ID:          "panic-recover",
	Category:    "Error Handling",
	Name:        "81. Panic and Recover",
	Description: "Catch panic with recover",
	Instruction: "Use defer with an anonymous function that calls recover and prints the word recovered if the recovered value is not nil, then call panic with the message oops",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered")
		}
	}()
	panic("oops")
}`,
	ExpectedOutput: "recovered",
	Difficulty:     "beginner",
	Explanation:    "panic stops normal execution and begins unwinding the stack, running deferred functions. recover() catches a panic, stopping the unwinding and returning the panic value. recover only works inside deferred functions. Use panic/recover for truly exceptional situations, not normal error handling.",
	Example:        "func safeDivide(a, b int) (result int) {\n    defer func() {\n        if r := recover(); r != nil {\n            fmt.Println(\"recovered from:\", r)\n            result = 0\n        }\n    }()\n    return a / b  // panics if b == 0\n}",
	UseCase:        "Use panic for unrecoverable errors (programmer mistakes, impossible situations). Use recover in servers or long-running programs to prevent crashes. Prefer returning errors for expected failures. Don't use panic/recover for control flow.",
	Prerequisites:  []string{"defer"},
	RelatedTopics:  []string{"defer", "error-check"},
	DocsURL:        "https://go.dev/blog/defer-panic-and-recover",
}

func init() {
	Register(Concept081)
}
