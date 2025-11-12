package concepts

// 073. json.Unmarshal
var Concept073 = Concept{
	Number:      73,
	ID:          "json-unmarshal",
	Category:    "Standard Library",
	Name:        "73. json.Unmarshal",
	Description: "Decode JSON to struct",
	Instruction: "Define type T as a struct with field X of type int, create a variable t of type T, use the Unmarshal function from the json package to parse JSON with X set to 2 into a pointer to t while ignoring errors, then print the X field of t",
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
	var t T
	json.Unmarshal([]byte(`,
	ExpectedOutput: "2",
	Difficulty:     "beginner",
	Explanation:    "json.Unmarshal parses JSON-encoded data and stores it in a Go value. Pass a pointer to the destination variable. JSON field names are matched to struct fields (case-insensitive). Returns an error if parsing fails.",
	Example:        "type Person struct {\n    Name string\n    Age  int\n}\ndata := []byte(`{\"name\":\"Bob\",\"age\":25}`)\nvar p Person\nerr := json.Unmarshal(data, &p)  // p.Name=\"Bob\", p.Age=25\nif err != nil {\n    log.Fatal(err)\n}",
	UseCase:        "Use json.Unmarshal to parse JSON from APIs, config files, or storage into Go structs. For streaming/large JSON, use json.Decoder instead. Use struct tags to map JSON field names to struct fields.",
	Prerequisites:  []string{"struct-declaration", "pointer-basics"},
	RelatedTopics:  []string{"json-marshal"},
	DocsURL:        "https://pkg.go.dev/encoding/json#Unmarshal",
}

func init() {
	Register(Concept073)
}
