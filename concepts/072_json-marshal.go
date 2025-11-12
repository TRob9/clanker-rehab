package concepts

// 072. json.Marshal
var Concept072 = Concept{
	Number:      72,
	ID:          "json-marshal",
	Category:    "Standard Library",
	Name:        "72. json.Marshal",
	Description: "Encode struct to JSON",
	Instruction: "Define type T as a struct with field X of type int, use the Marshal function from the json package to convert an instance with X set to 1 to JSON bytes while ignoring errors, convert to string, then print",
	Boilerplate: `package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// Your code here
}`,
	Answer: `package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type T struct { X int }
	data, _ := json.Marshal(T{X: 1})
	fmt.Println(string(data))
}`,
	ExpectedOutput: "{\"X\":1}",
	Difficulty:     "beginner",
	Explanation:    "json.Marshal converts Go values (structs, maps, slices) to JSON-encoded byte slices. Only exported (capitalized) struct fields are marshaled. Returns the JSON bytes and an error.",
	Example:        "type Person struct {\n    Name string\n    Age  int\n}\np := Person{\"Alice\", 30}\ndata, err := json.Marshal(p)  // {\"Name\":\"Alice\",\"Age\":30}\nif err != nil {\n    log.Fatal(err)\n}\nfmt.Println(string(data))",
	UseCase:        "Use json.Marshal to serialize Go data for APIs, config files, logging, or storage. For custom JSON field names, use struct tags: `json:\"name\"`. For pretty-printed JSON, use json.MarshalIndent.",
	Prerequisites:  []string{"struct-declaration"},
	RelatedTopics:  []string{"json-unmarshal"},
	DocsURL:        "https://pkg.go.dev/encoding/json#Marshal",
}

func init() {
	Register(Concept072)
}
