package concepts

// 079. errors.Is
var Concept079 = Concept{
	Number:      79,
	ID:          "errors-is",
	Category:    "Error Handling",
	Name:        "79. errors.Is",
	Description: "Check error identity",
	Instruction: "Create a base error using the New function from the errors package with message base, wrap it using the Errorf function with format "w: %w" and base, then print the result of using the Is function from the errors package to check if wrapped matches base",
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
	base := errors.New("base")
	wrapped := fmt.Errorf("w: %w", base)
	fmt.Println(errors.Is(wrapped, base))
}`,
	ExpectedOutput: "true",
	Difficulty:     "intermediate",
	Explanation:    "errors.Is checks if an error matches a target error, unwrapping the error chain as needed. It returns true if any error in the chain matches the target. Use this instead of direct error comparison (==) with wrapped errors.",
	Example:        "var ErrNotFound = errors.New(\"not found\")\n\nfunc getUser(id int) error {\n    return fmt.Errorf(\"user %d: %w\", id, ErrNotFound)\n}\n\nerr := getUser(123)\nif errors.Is(err, ErrNotFound) {\n    // Handle not found case\n}",
	UseCase:        "Use errors.Is to check for specific sentinel errors in wrapped error chains. Essential when using error wrapping - direct comparison (==) won't work through wrapping. Define sentinel errors as package-level variables for comparison.",
	Prerequisites:  []string{"error-wrapping"},
	RelatedTopics:  []string{"error-wrapping", "errors-as"},
	DocsURL:        "https://pkg.go.dev/errors#Is",
}

func init() {
	Register(Concept079)
}
