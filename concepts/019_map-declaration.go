package concepts

// 019. Map Declaration
var Concept019 = Concept{
	Number:      19,
	ID:          "map-declaration",
	Category:    "Data Structures",
	Name:        "19. Map Declaration",
	Description: "Create map with literal",
	Instruction: "Create a map with string keys and integer values, with key a having value 1, then print the value for key a",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	m := map[string]int{"a": 1}
	fmt.Println(m["a"])
}`,
	ExpectedOutput: "1",
	Difficulty:     "beginner",
	Explanation:    "Maps are Go's built-in hash table type, mapping keys to values. Maps are reference types and must be initialized before use (using make or a map literal). The zero value of a map is nil.",
	Example:        "m1 := map[string]int{\"a\": 1, \"b\": 2}\nm2 := make(map[string]int)\nvar m3 map[string]int  // nil map, can't add to it\n// Composite key:\nm4 := map[[2]int]string{{1,2}: \"point\"}",
	UseCase:        "Use maps for lookups, caches, counting, grouping, and any key-value storage. Maps provide O(1) average-case lookup and are the standard way to implement dictionaries, sets (map[T]bool), and associative arrays in Go.",
	Prerequisites:  []string{"var-declaration"},
	RelatedTopics:  []string{"map-make", "map-existence", "map-delete"},
	DocsURL:        "https://go.dev/tour/moretypes/19",
}

func init() {
	Register(Concept019)
}
