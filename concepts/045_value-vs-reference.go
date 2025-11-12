package concepts

// 045. Pass by Value vs Reference
var Concept045 = Concept{
	Number:      45,
	ID:          "value-vs-reference",
	Category:    "Pointers & Methods",
	Name:        "45. Pass by Value vs Reference",
	Description: "Function with pointer arg can mutate",
	Instruction: "Define a function named inc with parameter x as a pointer to int that increments the value at x, create variable n with value 1, call inc passing the address of n, then print n",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func inc(x *int) { *x++ }

func main() {
	n := 1
	inc(&n)
	fmt.Println(n)
}`,
	ExpectedOutput: "2",
	Difficulty:     "beginner",
	Explanation:    "Go is always pass-by-value, but you can pass a pointer value to allow functions to modify the original. Passing a pointer copies the pointer (not the data it points to), enabling mutation of the original value.",
	Example:        "// Pass by value (copy):\nfunc double(x int) {\n    x = x * 2  // only modifies copy\n}\nn := 5\ndouble(n)\n// n is still 5\n\n// Pass pointer (can mutate):\nfunc doublePtr(x *int) {\n    *x = *x * 2  // modifies original\n}\nm := 5\ndoublePtr(&m)\n// m is now 10",
	UseCase:        "Pass pointers to functions when you need to modify arguments, avoid copying large structs, or return multiple values via output parameters. Common pattern: func Parse(input string, result *Data) error.",
	Prerequisites:  []string{"pointer-basics", "dereference"},
	RelatedTopics:  []string{"pointer-basics", "method-pointer-receiver", "pointer-mutation"},
	DocsURL:        "https://go.dev/doc/faq#pass_by_value",
}

func init() {
	Register(Concept045)
}
