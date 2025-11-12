package concepts

// 082. Package Declaration
var Concept082 = Concept{
	Number:      82,
	ID:          "package-declaration",
	Category:    "Tooling & Tests",
	Name:        "82. Package Declaration",
	Description: "Declare package name",
	Instruction: "Add the package declaration for main at the top",
	Boilerplate: `// Your code here

import "fmt"

func main() {
	fmt.Println("test")
}`,
	Answer: `package main

import "fmt"

func main() {
	fmt.Println("test")
}`,
	ExpectedOutput: "test",
	Difficulty:     "beginner",
	Explanation:    "Every Go file starts with a package declaration. package main creates an executable program with a main() function as entry point. Other package names create libraries. Package names should be lowercase, short, and descriptive.",
	Example:        "// Executable:\npackage main\nfunc main() { ... }\n\n// Library:\npackage utils\nfunc Helper() { ... }\n\n// Test file:\npackage utils_test  // external test\nfunc TestHelper(t *testing.T) { ... }",
	UseCase:        "Use package main for executables. Use descriptive names for libraries (http, json, utils). Test packages can use either the same name (white-box testing) or package_test (black-box testing). Package names affect import paths.",
	Prerequisites:  nil,
	RelatedTopics:  nil,
	DocsURL:        "https://go.dev/doc/effective_go#package-names",
}

func init() {
	Register(Concept082)
}
