package concepts

// 069. strconv.Atoi
var Concept069 = Concept{
	Number:      69,
	ID:          "strconv-atoi",
	Category:    "Standard Library",
	Name:        "69. strconv.Atoi",
	Description: "Convert string to int",
	Instruction: "Use the Atoi function from the strconv package to convert the string "123" to an integer while ignoring any error, then print the result",
	Boilerplate: `package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Your code here (ignore error for simplicity)
}`,
	Answer: `package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, _ := strconv.Atoi("123")
	fmt.Println(i)
}`,
	ExpectedOutput: "123",
	Difficulty:     "beginner",
	Explanation:    "strconv.Atoi (ASCII to integer) converts a string to an int. It returns the integer and an error (nil on success). The string must contain a valid decimal integer, optionally preceded by a sign.",
	Example:        "i, err := strconv.Atoi(\"123\")   // i=123, err=nil\ni, err = strconv.Atoi(\"-456\")     // i=-456, err=nil\ni, err = strconv.Atoi(\"abc\")      // i=0, err!=nil\n// Always check error:\nif err != nil {\n    log.Fatal(err)\n}",
	UseCase:        "Use strconv.Atoi to parse user input, read numeric config values, or process numeric strings from files/APIs. Always check the error return value. For base 10 int conversion with explicit bit size, use strconv.ParseInt.",
	Prerequisites:  []string{"multiple-return", "error-check"},
	RelatedTopics:  []string{"strconv-itoa", "type-conversion"},
	DocsURL:        "https://pkg.go.dev/strconv#Atoi",
}

func init() {
	Register(Concept069)
}
