package concepts

// 085. Init Function
var Concept085 = Concept{
	Number:      85,
	ID:          "init-function",
	Category:    "Miscellaneous",
	Name:        "85. Init Function",
	Description: "Run code before main with init()",
	Instruction: "Declare a package-level variable x as an integer, create an init function that sets x to 10, then print x in the main function",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

var x int

func init() {
	x = 10
}

func main() {
	fmt.Println(x)
}`,
	ExpectedOutput: "10",
	Difficulty:     "beginner",
	Explanation:    "The init() function runs automatically before main(), used for package initialization. Multiple init() functions can exist in a file/package and run in declaration order. init() takes no parameters and returns nothing.",
	Example:        "var config map[string]string\n\nfunc init() {\n    config = make(map[string]string)\n    config[\"env\"] = \"dev\"\n}\n\nfunc main() {\n    // config is already initialized\n    fmt.Println(config[\"env\"])\n}",
	UseCase:        "Use init() for package-level setup: initializing complex globals, registering drivers, validating configuration, or one-time setup. Runs before main() in executables or before package use in libraries. Keep init() simple and avoid side effects when possible.",
	Prerequisites:  nil,
	RelatedTopics:  []string{"blank-import"},
	DocsURL:        "https://go.dev/doc/effective_go#init",
}

func init() {
	Register(Concept085)
}
