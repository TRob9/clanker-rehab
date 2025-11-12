package concepts

// 074. errors.New
var Concept074 = Concept{
	Number:      74,
	ID:          "errors-new",
	Category:    "Standard Library",
	Name:        "74. errors.New",
	Description: "Create simple error",
	Instruction: "Use errors.New to create an error with message \"fail\", assign to err, then print err.Error()",
	Boilerplate: `package main

import (
	"errors"
	"fmt"
)

func main() {
	// Your code here
}`,
	Answer: `package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("fail")
	fmt.Println(err.Error())
}`,
	ExpectedOutput: "fail",
	Difficulty:     "beginner",
	Explanation:    "errors.New creates a simple error with a static string message. The error interface requires only an Error() string method. This is the simplest way to create errors in Go.",
	Example:        "err := errors.New(\"something went wrong\")\nif err != nil {\n    fmt.Println(err.Error())  // \"something went wrong\"\n}\n// Common pattern:\nfunc doSomething() error {\n    return errors.New(\"failed\")\n}",
	UseCase:        "Use errors.New for simple, static error messages. For formatted errors with dynamic values, use fmt.Errorf. For custom error types with additional fields, define your own type implementing the error interface.",
	Prerequisites:  nil,
	RelatedTopics:  []string{"error-interface", "custom-error", "error-wrapping"},
	DocsURL:        "https://pkg.go.dev/errors#New",
}

func init() {
	Register(Concept074)
}
