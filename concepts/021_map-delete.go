package concepts

// 021. Delete from Map
var Concept021 = Concept{
	Number:      21,
	ID:          "map-delete",
	Category:    "Data Structures",
	Name:        "21. Delete from Map",
	Description: "Use delete to remove key",
	Instruction: "Create a map with string keys and integer values containing key a with value 1 and key b with value 2, delete key a from the map, then print the length of the map",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	m := map[string]int{"a": 1, "b": 2}
	delete(m, "a")
	fmt.Println(len(m))
}`,
	ExpectedOutput: "1",
	Difficulty:     "beginner",
	Explanation:    "The delete() built-in function removes a key-value pair from a map. If the key doesn't exist, delete does nothing (no error). After deletion, accessing the key returns the zero value.",
	Example:        "m := map[string]int{\"a\": 1, \"b\": 2}\ndelete(m, \"a\")\ndelete(m, \"nonexistent\")  // safe, does nothing\n// Check before delete:\nif _, exists := m[\"b\"]; exists {\n    delete(m, \"b\")\n}",
	UseCase:        "Use delete to remove unwanted entries, implement cache eviction, clean up temporary data, or maintain a limited-size map. delete is safe even if the key doesn't exist.",
	Prerequisites:  []string{"map-declaration"},
	RelatedTopics:  []string{"map-declaration", "map-existence"},
	DocsURL:        "https://go.dev/tour/moretypes/22",
}

func init() {
	Register(Concept021)
}
