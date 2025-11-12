package concepts

// 077. Custom Error Type
var Concept077 = Concept{
	Number:      77,
	ID:          "custom-error",
	Category:    "Error Handling",
	Name:        "77. Custom Error Type",
	Description: "Define custom error",
	Instruction: "Define type MyError as a struct with Msg field of type string, implement the Error method returning the Msg field, create an error as MyError with Msg set to the word custom, then print the result of calling the Error method",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

type MyError struct { Msg string }

func (e MyError) Error() string { return e.Msg }

func main() {
	err := MyError{"custom"}
	fmt.Println(err.Error())
}`,
	ExpectedOutput: "custom",
	Difficulty:     "beginner",
	Explanation:    "Create custom error types by defining a struct with an Error() string method. This lets you add extra fields (error codes, context, timestamps) beyond just a message. Custom errors enable type-based error handling with errors.As().",
	Example:        "type ValidationError struct {\n    Field string\n    Value string\n}\n\nfunc (e ValidationError) Error() string {\n    return fmt.Sprintf(\"invalid %s: %s\", e.Field, e.Value)\n}\n\nfunc validate(age int) error {\n    if age < 0 {\n        return ValidationError{\"age\", strconv.Itoa(age)}\n    }\n    return nil\n}",
	UseCase:        "Use custom errors when you need to carry additional context, enable error type checking, implement retry logic, or build error hierarchies. Much more powerful than plain string errors. Essential for library APIs.",
	Prerequisites:  []string{"error-interface", "struct-declaration"},
	RelatedTopics:  []string{"error-interface", "errors-as"},
	DocsURL:        "https://go.dev/blog/error-handling-and-go",
}

func init() {
	Register(Concept077)
}
