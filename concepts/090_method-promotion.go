package concepts

// 090. Method Promotion via Embedding
var Concept090 = Concept{
	Number:      90,
	ID:          "method-promotion",
	Category:    "Miscellaneous",
	Name:        "90. Method Promotion via Embedding",
	Description: "Embedded type's methods are promoted",
	Instruction: "Define type A with a method M that returns nine, define type B that embeds A, create an instance b of type B, then call and print the promoted method M on b",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

type A struct{}

func (A) M() int { return 9 }

type B struct { A }

func main() {
	b := B{}
	fmt.Println(b.M())
}`,
	ExpectedOutput: "9",
	Difficulty:     "beginner",
	Explanation:    "When a struct embeds another type, the embedded type's methods are promoted to the outer struct. This means you can call the embedded type's methods directly on the outer struct, enabling easy composition and delegation.",
	Example:        "type Logger struct {}\nfunc (l Logger) Log(msg string) {\n    fmt.Println(\"LOG:\", msg)\n}\n\ntype Service struct {\n    Logger  // embedded - methods promoted\n    name string\n}\n\ns := Service{name: \"API\"}\ns.Log(\"started\")  // Calls Logger.Log directly\n// Same as s.Logger.Log(\"started\")",
	UseCase:        "Use method promotion to add functionality to types through composition, create mixins, or satisfy interfaces by embedding types that already implement them. This is Go's primary mechanism for code reuse and polymorphism.",
	Prerequisites:  []string{"struct-embedding", "method-value-receiver"},
	RelatedTopics:  []string{"struct-embedding", "interface-basic", "embedding-override"},
	DocsURL:        "https://go.dev/doc/effective_go#embedding",
}

func init() {
	Register(Concept090)
}
