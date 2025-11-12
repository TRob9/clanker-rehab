package concepts

// 068. strings.Join
var Concept068 = Concept{
	Number:      68,
	ID:          "strings-join",
	Category:    "Standard Library",
	Name:        "68. strings.Join",
	Description: "Join slice into string",
	Instruction: "Use the Join function from the strings package to join a slice containing a and b with hyphen separator, then print the result",
	Boilerplate: `package main

import (
	"fmt"
	"strings"
)

func main() {
	// Your code here
}`,
	Answer: `package main

import (
	"fmt"
	"strings"
)

func main() {
	result := strings.Join([]string{"a", "b"}, "-")
	fmt.Println(result)
}`,
	ExpectedOutput: "a-b",
	Difficulty:     "beginner",
	Explanation:    "strings.Join concatenates the elements of a string slice into a single string, separated by a delimiter. It's the inverse operation of strings.Split. More efficient than manual concatenation with + for multiple strings.",
	Example:        "strings.Join([]string{\"a\", \"b\", \"c\"}, \",\")  // \"a,b,c\"\nstrings.Join([]string{\"go\", \"is\", \"fun\"}, \" \") // \"go is fun\"\nstrings.Join([]string{\"x\", \"y\"}, \"\")        // \"xy\"\nstrings.Join([]string{\"single\"}, \",\")      // \"single\"",
	UseCase:        "Use strings.Join to build CSV lines, create paths, format lists for display, or combine string slices efficiently. Much faster than repeated string concatenation with + or += for multiple strings.",
	Prerequisites:  []string{"slice-basics"},
	RelatedTopics:  []string{"strings-split"},
	DocsURL:        "https://pkg.go.dev/strings#Join",
}

func init() {
	Register(Concept068)
}
