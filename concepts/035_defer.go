package concepts

// 035. Defer Statement
var Concept035 = Concept{
	Number:      35,
	ID:          "defer",
	Category:    "Functions & Closures",
	Name:        "35. Defer Statement",
	Description: "Defer execution until function returns",
	Instruction: "Use defer to print "world", then print "hello" with a trailing space, which should output hello world",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	defer fmt.Print("world")
	fmt.Print("hello ")
}`,
	ExpectedOutput: "hello world",
	Difficulty:     "beginner",
	Explanation:    "The defer keyword schedules a function call to execute after the surrounding function returns. Defer is commonly used for cleanup tasks like closing files, unlocking mutexes, or rolling back transactions. Arguments are evaluated immediately, but the call is delayed.",
	Example:        "func readFile(path string) error {\n    f, err := os.Open(path)\n    if err != nil { return err }\n    defer f.Close()  // ensures file is closed\n    \n    // use file...\n    return nil\n}\n// Multiple defers stack (LIFO)",
	UseCase:        "Use defer for cleanup code that must run regardless of how the function exits (return, panic, etc.). Perfect for resource cleanup, unlocking mutexes, and ensuring state is restored. Defers keep cleanup code near allocation code.",
	Prerequisites:  []string{"function-params"},
	RelatedTopics:  []string{"defer-order", "panic-recover", "named-return"},
	DocsURL:        "https://go.dev/tour/flowcontrol/12",
}

func init() {
	Register(Concept035)
}
