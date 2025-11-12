package concepts

// 046. Interface Definition
var Concept046 = Concept{
	Number:      46,
	ID:          "interface-basic",
	Category:    "Interfaces",
	Name:        "46. Interface Definition",
	Description: "Define and implement interface",
	Instruction: "Implement the Speak method on the Dog type to return the word woof",
	Boilerplate: `package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	// Your code here
}

func main() {
	var s Speaker = Dog{}
	fmt.Println(s.Speak())
}`,
	Answer: `package main

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "woof"
}

func main() {
	var s Speaker = Dog{}
	fmt.Println(s.Speak())
}`,
	ExpectedOutput: "woof",
	TestCases: []TestCase{
		{Input: "Dog{}.Speak()", Expected: "woof"},
	},
	Difficulty:     "beginner",
	Explanation:    "Interfaces define behavior through method signatures. Types implement interfaces implicitly by implementing all required methods - no explicit 'implements' declaration needed. This enables polymorphism and decoupling in Go.",
	Example:        "type Writer interface {\n    Write([]byte) (int, error)\n}\n\ntype File struct{}\nfunc (f File) Write(data []byte) (int, error) {\n    // File automatically implements Writer\n    return len(data), nil\n}\n\nvar w Writer = File{}",
	UseCase:        "Use interfaces to define contracts, enable polymorphism, accept multiple types, and decouple code. Interfaces are central to Go's design philosophy. Keep interfaces small - often just one or two methods.",
	Prerequisites:  []string{"method-value-receiver"},
	RelatedTopics:  []string{"empty-interface", "type-assertion", "interface-embedding"},
	DocsURL:        "https://go.dev/tour/methods/9",
}

func init() {
	Register(Concept046)
}
