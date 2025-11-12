package concepts

// 051. Interface Embedding
var Concept051 = Concept{
	Number:      51,
	ID:          "interface-embedding",
	Category:    "Interfaces",
	Name:        "51. Interface Embedding",
	Description: "Compose interfaces",
	Instruction: "Define a Reader interface with a Read method that returns a string. Define a Writer interface with a Write method that returns a string. Define a ReadWriter interface that embeds both. Implement a type T with both methods where Read returns the word data and Write returns an empty string, create a ReadWriter variable, and print the result of calling Read.",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

type Reader interface { Read() string }
type Writer interface { Write() string }
type ReadWriter interface { Reader; Writer }
type T struct{}

func (T) Read() string { return "data" }
func (T) Write() string { return "" }

func main() {
	var rw ReadWriter = T{}
	fmt.Println(rw.Read())
}`,
	ExpectedOutput: "data",
	Difficulty:     "intermediate",
	Explanation:    "Interfaces can embed other interfaces, combining their method sets. A type must implement all methods from all embedded interfaces to satisfy the composite interface. This enables building complex interfaces from simple ones.",
	Example:        "type Reader interface {\n    Read([]byte) (int, error)\n}\n\ntype Writer interface {\n    Write([]byte) (int, error)\n}\n\ntype ReadWriter interface {\n    Reader  // embeds Reader\n    Writer  // embeds Writer\n}\n// ReadWriter requires both Read and Write methods",
	UseCase:        "Use interface embedding to compose interfaces, build layered abstractions, and avoid repeating method signatures. The io package uses this heavily: io.ReadWriter embeds io.Reader and io.Writer.",
	Prerequisites:  []string{"interface-basic"},
	RelatedTopics:  []string{"interface-basic", "multiple-interfaces"},
	DocsURL:        "https://go.dev/doc/effective_go#embedding",
}

func init() {
	Register(Concept051)
}
