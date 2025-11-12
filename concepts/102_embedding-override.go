package concepts

// 102. Method Override via Embedding
var Concept102 = Concept{
	Number:      102,
	ID:          "embedding-override",
	Category:    "Pointers & Methods",
	Name:        "102. Method Override via Embedding",
	Description: "Override embedded type's method in outer type",
	Instruction: "Define type A with method M returning the letter A. Embed A in type B and override method M to return the letter B. Create an instance of B and print the result of calling method M.",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

type A struct{}

func (A) M() string { return "A" }

type B struct { A }

func (B) M() string { return "B" }

func main() {
	b := B{}
	fmt.Println(b.M())
}`,
	ExpectedOutput: "B",
	Difficulty:     "advanced",
	Explanation:    "When an outer type defines a method with the same name as an embedded type's method, the outer method takes precedence (shadows the inner one). You can still access the embedded type's method via the field name. This is Go's way of method overriding.",
	Example:        "type Base struct {}\nfunc (b Base) Greet() string { return \"Base\" }\n\ntype Derived struct { Base }\nfunc (d Derived) Greet() string { return \"Derived\" }\n\nd := Derived{}\nd.Greet()       // \"Derived\" (override)\nd.Base.Greet()  // \"Base\" (original)",
	UseCase:        "Use method overriding to customize behavior in composed types, implement specializations, or adapt embedded types. Common in decorator patterns and when extending third-party types. Remember: Go favors composition over inheritance.",
	Prerequisites:  []string{"struct-embedding", "method-value-receiver"},
	RelatedTopics:  []string{"struct-embedding", "method-promotion"},
	DocsURL:        "https://go.dev/doc/effective_go#embedding",
}

func init() {
	Register(Concept102)
}
