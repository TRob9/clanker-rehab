package concepts

// 050. Multiple Interface Implementation
var Concept050 = Concept{
	Number:      50,
	ID:          "multiple-interfaces",
	Category:    "Interfaces",
	Name:        "50. Multiple Interface Implementation",
	Description: "Type can satisfy multiple interfaces",
	Instruction: "Define two interfaces A and B with methods M1 and M2 that return integers. Create a type T that implements both methods returning 1 and 2. Create a variable of type A with value T and print the result of calling M1",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

type A interface { M1() int }
type B interface { M2() int }
type T struct{}

func (T) M1() int { return 1 }
func (T) M2() int { return 2 }

func main() {
	var a A = T{}
	fmt.Println(a.M1())
}`,
	ExpectedOutput: "1",
	Difficulty:     "intermediate",
	Explanation:    "A single type can implement multiple interfaces by implementing all required methods from each interface. This enables flexible composition and allows types to satisfy different contracts simultaneously.",
	Example:        "type Reader interface { Read() string }\ntype Writer interface { Write(string) }\n\ntype File struct{}\nfunc (f File) Read() string { return \"data\" }\nfunc (f File) Write(s string) { }\n\n// File implements both Reader and Writer\nvar r Reader = File{}\nvar w Writer = File{}",
	UseCase:        "Use multiple interfaces to compose behavior, create types that fit multiple contracts, and build flexible APIs. Common pattern: io.ReadWriter combines io.Reader and io.Writer.",
	Prerequisites:  []string{"interface-basic"},
	RelatedTopics:  []string{"interface-basic", "interface-embedding"},
	DocsURL:        "https://go.dev/tour/methods/9",
}

func init() {
	Register(Concept050)
}
