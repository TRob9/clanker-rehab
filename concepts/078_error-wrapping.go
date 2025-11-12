package concepts

// 078. Error Wrapping
var Concept078 = Concept{
	Number:      78,
	ID:          "error-wrapping",
	Category:    "Error Handling",
	Name:        "78. Error Wrapping",
	Description: "Wrap error with fmt.Errorf %w",
	Instruction: "Create an error using the New function from the errors package with message base, wrap it using the Errorf function from the fmt package with format string wrap: %w and the error, then print the result of calling the Error method on the wrapped error",
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
	err := errors.New("base")
	wrapped := fmt.Errorf("wrap: %w", err)
	fmt.Println(wrapped.Error())
}`,
	ExpectedOutput: "wrap: base",
	Difficulty:     "beginner",
	Explanation:    "Error wrapping adds context to errors while preserving the original error. Use fmt.Errorf with %w verb to wrap an error. The wrapped error can be unwrapped with errors.Unwrap, errors.Is, or errors.As.",
	Example:        "func readConfig() error {\n    err := os.ReadFile(\"config.json\")\n    if err != nil {\n        return fmt.Errorf(\"failed to read config: %w\", err)\n    }\n    return nil\n}\n// Creates error chain: \"failed to read config: file not found\"\n// errors.Is can check the original error",
	UseCase:        "Wrap errors to add context as they bubble up the call stack. This preserves the error chain while adding information at each layer. Essential for debugging - you get both the root cause and the context of where it occurred.",
	Prerequisites:  []string{"errors-new", "error-check"},
	RelatedTopics:  []string{"errors-is", "errors-as", "errors-new"},
	DocsURL:        "https://go.dev/blog/go1.13-errors",
}

func init() {
	Register(Concept078)
}
