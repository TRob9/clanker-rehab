package concepts

// 080. errors.As
var Concept080 = Concept{
	Number:      80,
	ID:          "errors-as",
	Category:    "Error Handling",
	Name:        "80. errors.As",
	Description: "Extract error type",
	Instruction: "Define type MyErr as a struct with Code field of type int, implement the Error method returning the letter e, create an error as MyErr with Code set to 5, use the As function from the errors package to extract the error into a variable, then print the Code field",
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

type MyErr struct { Code int }

func (e MyErr) Error() string { return "e" }

func main() {
	err := MyErr{Code: 5}
	var me MyErr
	errors.As(err, &me)
	fmt.Println(me.Code)
}`,
	ExpectedOutput: "5",
	Difficulty:     "intermediate",
	Explanation:    "errors.As finds the first error in the chain that matches the target type and assigns it to the target. Returns true if a match is found. Use this to extract custom error types from wrapped errors to access their fields.",
	Example:        "type ValidationError struct {\n    Field string\n}\n\nfunc (e ValidationError) Error() string { return e.Field }\n\nerr := fmt.Errorf(\"validation failed: %w\", ValidationError{\"email\"})\n\nvar ve ValidationError\nif errors.As(err, &ve) {\n    fmt.Println(\"invalid field:\", ve.Field)  // email\n}",
	UseCase:        "Use errors.As to extract custom error types from error chains and access their fields (error codes, retry info, context). Essential for type-based error handling with custom errors. Pass a pointer to your error type variable.",
	Prerequisites:  []string{"custom-error", "error-wrapping"},
	RelatedTopics:  []string{"custom-error", "errors-is", "error-wrapping"},
	DocsURL:        "https://pkg.go.dev/errors#As",
}

func init() {
	Register(Concept080)
}
