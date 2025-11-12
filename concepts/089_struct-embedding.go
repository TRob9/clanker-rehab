package concepts

// 089. Struct Embedding
var Concept089 = Concept{
	Number:      89,
	ID:          "struct-embedding",
	Category:    "Miscellaneous",
	Name:        "89. Struct Embedding",
	Description: "Embed struct for composition",
	Instruction: "Define type A with field X of type int, define type B that embeds A, create an instance b with the embedded X field set to 3, then print the promoted X field of b",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

type A struct { X int }
type B struct { A }

func main() {
	b := B{A: A{X: 3}}
	fmt.Println(b.X)
}`,
	ExpectedOutput: "3",
	Difficulty:     "beginner",
	Explanation:    "Struct embedding allows one struct to include another struct without a field name, promoting the embedded struct's fields and methods to the outer struct. This enables composition and field/method promotion in Go's type system.",
	Example:        "type Base struct {\n    ID int\n    Name string\n}\n\ntype User struct {\n    Base  // embedded - no field name\n    Email string\n}\n\nu := User{\n    Base: Base{ID: 1, Name: \"Alice\"},\n    Email: \"alice@example.com\",\n}\nfmt.Println(u.ID)    // Access promoted field directly\nfmt.Println(u.Name)  // No need for u.Base.Name",
	UseCase:        "Use struct embedding for composition over inheritance, to extend types without modifying them, or to promote fields/methods from inner types. Common in implementing interfaces by embedding types that already implement them.",
	Prerequisites:  []string{"struct-declaration", "nested-struct"},
	RelatedTopics:  []string{"method-promotion", "interface-embedding", "nested-struct"},
	DocsURL:        "https://go.dev/doc/effective_go#embedding",
}

func init() {
	Register(Concept089)
}
