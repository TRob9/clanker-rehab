package concepts

// 022. Check Map Key Existence
var Concept022 = Concept{
	Number:      22,
	ID:          "map-existence",
	Category:    "Data Structures",
	Name:        "22. Check Map Key Existence",
	Description: "Use comma ok idiom",
	Instruction: "Create a map with string keys and integer values containing key a with value 1, use the comma-ok idiom to check if key b exists by ignoring the value and capturing the boolean, then print the boolean result",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	m := map[string]int{"a": 1}
	_, ok := m["b"]
	fmt.Println(ok)
}`,
	ExpectedOutput: "false",
	Difficulty:     "beginner",
	Explanation:    "Check if a map key exists using the comma-ok idiom: value, ok := m[key]. If the key exists, ok is true and value is the associated value. If not, ok is false and value is the zero value of the value type.",
	Example:        "m := map[string]int{\"a\": 1}\nval, ok := m[\"a\"]    // val=1, ok=true\nval2, ok2 := m[\"b\"]  // val2=0, ok2=false\n// Common pattern:\nif score, exists := scores[name]; exists {\n    fmt.Println(score)\n}",
	UseCase:        "Use the comma-ok idiom to distinguish between 'key doesn't exist' and 'key exists with zero value'. This is crucial for maps with integer or boolean values where zero is meaningful.",
	Prerequisites:  []string{"map-declaration"},
	RelatedTopics:  []string{"map-declaration", "type-assertion-ok", "channel-close"},
	DocsURL:        "https://go.dev/tour/moretypes/22",
}

func init() {
	Register(Concept022)
}
