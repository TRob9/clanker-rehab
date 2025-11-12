package concepts

// 014. Slice Basics
var Concept014 = Concept{
	Number:      14,
	ID:          "slice-basics",
	Category:    "Data Structures",
	Name:        "14. Slice Basics",
	Description: "Create a slice with literal",
	Instruction: "Create a slice containing the integers 10 and 20, then print the first element",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	s := []int{10, 20}
	fmt.Println(s[0])
}`,
	ExpectedOutput: "10",
	Difficulty:     "beginner",
	Explanation:    "Slices are dynamically-sized, flexible views into arrays. Unlike arrays, slices are reference types. A slice doesn't store data itself - it's a descriptor pointing to an underlying array.",
	Example:        "s := []int{1, 2, 3, 4, 5}\n// Empty slice:\nvar empty []int\n// Slice from array:\narr := [5]int{1,2,3,4,5}\nslice := arr[1:4]  // [2, 3, 4]",
	UseCase:        "Slices are Go's primary collection type. Use them for lists, stacks, dynamic arrays, and whenever you need a growable sequence. They're the idiomatic way to handle collections in Go.",
	Prerequisites:  []string{"array-declaration"},
	RelatedTopics:  []string{"slice-make", "slice-append", "slice-slicing", "array-declaration"},
	DocsURL:        "https://go.dev/tour/moretypes/7",
}

func init() {
	Register(Concept014)
}
