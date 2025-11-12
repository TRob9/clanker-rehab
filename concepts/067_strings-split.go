package concepts

// 067. strings.Split
var Concept067 = Concept{
	Number:      67,
	ID:          "strings-split",
	Category:    "Standard Library",
	Name:        "67. strings.Split",
	Description: "Split string into slice",
	Instruction: "Use the Split function from the strings package to split the string a,b,c by comma, then print the length of the result",
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
	parts := strings.Split("a,b,c", ",")
	fmt.Println(len(parts))
}`,
	ExpectedOutput: "3",
	Difficulty:     "beginner",
	Explanation:    "strings.Split divides a string into a slice of substrings separated by a delimiter. The delimiter is not included in the results. If the delimiter is empty, it splits after each UTF-8 character.",
	Example:        "strings.Split(\"a,b,c\", \",\")   // [\"a\" \"b\" \"c\"]\nstrings.Split(\"a-b-c\", \"-\")   // [\"a\" \"b\" \"c\"]\nstrings.Split(\"hello\", \"\")    // [\"h\" \"e\" \"l\" \"l\" \"o\"]\nstrings.Split(\"no-sep\", \",\")  // [\"no-sep\"]",
	UseCase:        "Use strings.Split to parse CSV data, break apart paths, tokenize input, or extract fields from delimited text. For splitting with multiple delimiters or more complex patterns, use strings.FieldsFunc or regular expressions.",
	Prerequisites:  []string{"slice-basics"},
	RelatedTopics:  []string{"strings-join", "strings-contains"},
	DocsURL:        "https://pkg.go.dev/strings#Split",
}

func init() {
	Register(Concept067)
}
