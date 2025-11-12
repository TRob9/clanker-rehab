package concepts

// 053. Error Interface
var Concept053 = Concept{
	Number:      53,
	ID:          "error-interface",
	Category:    "Interfaces",
	Name:        "53. Error Interface",
	Description: "Implement error interface",
	Instruction: "Define a MyErr struct with a Msg field. Implement the Error method that returns the Msg field. Create an error variable with MyErr initialized with message fail and print the result of calling the Error method",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

type MyErr struct { Msg string }

func (e MyErr) Error() string { return e.Msg }

func main() {
	err := MyErr{"fail"}
	fmt.Println(err.Error())
}`,
	ExpectedOutput: "fail",
	Difficulty:     "beginner",
	Explanation:    "The error interface has one method: Error() string. Any type implementing this method can be used as an error. This enables custom error types with additional context, fields, or behavior.",
	Example:        "type ValidationError struct {\n    Field string\n    Message string\n}\n\nfunc (e ValidationError) Error() string {\n    return fmt.Sprintf(\"%s: %s\", e.Field, e.Message)\n}\n\nfunc validate(age int) error {\n    if age < 0 {\n        return ValidationError{\"age\", \"must be positive\"}\n    }\n    return nil\n}",
	UseCase:        "Implement custom error types to carry additional context (error codes, stack traces, retry info), enable error type checking with errors.As(), and build rich error hierarchies. Much better than plain strings.",
	Prerequisites:  []string{"interface-basic"},
	RelatedTopics:  []string{"custom-error", "errors-as", "errors-is"},
	DocsURL:        "https://go.dev/blog/error-handling-and-go",
}

func init() {
	Register(Concept053)
}
