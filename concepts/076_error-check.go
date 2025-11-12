package concepts

// 076. Basic Error Check
var Concept076 = Concept{
	Number:      76,
	ID:          "error-check",
	Category:    "Error Handling",
	Name:        "76. Basic Error Check",
	Description: "Check if error is nil",
	Instruction: "Define a function f that returns an error and returns nil, call f and check if the error is not nil, print "err" if true or the word "ok" if false",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func f() error { return nil }

func main() {
	err := f()
	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println("ok")
	}
}`,
	ExpectedOutput: "ok",
	Difficulty:     "beginner",
	Explanation:    "Go uses explicit error handling - functions return error as the last return value. Check if error is nil to detect success. nil error means success, non-nil means failure. This is idiomatic Go error handling.",
	Example:        "func divide(a, b float64) (float64, error) {\n    if b == 0 {\n        return 0, errors.New(\"division by zero\")\n    }\n    return a / b, nil\n}\n\nresult, err := divide(10, 2)\nif err != nil {\n    log.Fatal(err)  // handle error\n}\nfmt.Println(result)  // use result",
	UseCase:        "Always check errors from functions that return them. Ignoring errors is a common source of bugs. Use if err != nil pattern immediately after function calls. For testing, you can ignore errors with _ but only in examples/tests.",
	Prerequisites:  []string{"multiple-return"},
	RelatedTopics:  []string{"errors-new", "error-wrapping"},
	DocsURL:        "https://go.dev/blog/error-handling-and-go",
}

func init() {
	Register(Concept076)
}
