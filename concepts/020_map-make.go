package concepts

// 020. Map with Make
var Concept020 = Concept{
	Number:      20,
	ID:          "map-make",
	Category:    "Data Structures",
	Name:        "20. Map with Make",
	Description: "Use make to create map",
	Instruction: "Create a map using make with string keys and integer values, set key x to value 10, then print the value for key x",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["x"] = 10
	fmt.Println(m["x"])
}`,
	ExpectedOutput: "10",
	Difficulty:     "beginner",
	Explanation:    "The make() function initializes a map, allocating and initializing its internal hash table. Unlike map literals, make creates an empty map ready for insertion. You can optionally provide an initial capacity hint.",
	Example:        "m := make(map[string]int)\nm[\"one\"] = 1\nm[\"two\"] = 2\n// With capacity hint:\nbig := make(map[string]int, 1000)",
	UseCase:        "Use make when you need an empty map or want to preallocate capacity for performance. Providing a capacity hint can reduce reallocations when you know the approximate size. make is required for maps that will be populated later.",
	Prerequisites:  []string{"map-declaration"},
	RelatedTopics:  []string{"map-declaration", "slice-make"},
	DocsURL:        "https://go.dev/tour/moretypes/20",
}

func init() {
	Register(Concept020)
}
