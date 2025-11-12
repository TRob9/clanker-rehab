package concepts

// 095. Reflection TypeOf
var Concept095 = Concept{
	Number:      95,
	ID:          "reflect-type",
	Category:    "Miscellaneous",
	Name:        "95. Reflection TypeOf",
	Description: "Use reflect package to inspect types at runtime",
	Instruction: "Use the TypeOf function from the reflect package to get the type of the integer 42, then print the Kind as a string, which should be the word int.",
	Boilerplate: `package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Your code here
}`,
	Answer: `package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.TypeOf(42)
	fmt.Println(t.Kind())
}`,
	ExpectedOutput: "int",
	Difficulty:     "advanced",
	Explanation:    "The reflect package provides runtime reflection, allowing programs to inspect and manipulate objects of arbitrary types. reflect.TypeOf returns a Type describing the type of the value. Kind() returns the specific kind (int, string, struct, etc.).",
	Example:        "import \"reflect\"\n\nvar x int = 42\nt := reflect.TypeOf(x)\nfmt.Println(t.Kind())  // int\nfmt.Println(t.String()) // int\n\n// For structs:\ntype Person struct { Name string }\np := Person{\"Alice\"}\nt = reflect.TypeOf(p)\nfmt.Println(t.Kind())  // struct\nfmt.Println(t.NumField()) // 1",
	UseCase:        "Use reflection for serialization/deserialization (JSON, XML), building generic libraries, dependency injection, or when types are unknown at compile time. Warning: reflection is powerful but slow and can break type safety. Use only when necessary.",
	Prerequisites:  []string{"interface-basic"},
	RelatedTopics:  []string{"empty-interface", "type-assertion"},
	DocsURL:        "https://pkg.go.dev/reflect",
}

func init() {
	Register(Concept095)
}
