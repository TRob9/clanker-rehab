package concepts

// 017. Slice Slicing
var Concept017 = Concept{
	Number:      17,
	ID:          "slice-slicing",
	Category:    "Data Structures",
	Name:        "17. Slice Slicing",
	Description: "Create sub-slice with [start:end]",
	Instruction: "Create a slice containing 1, 2, 3, and 4, then create a sub-slice from index 1 up to but not including index 3, then print the first element of the sub-slice",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4}
	sub := s[1:3]
	fmt.Println(sub[0])
}`,
	ExpectedOutput: "2",
	Difficulty:     "beginner",
	Explanation:    "Slicing creates a new slice from an existing slice or array using [low:high] syntax. The result includes elements from low up to (but not including) high. Slices share the same underlying array - changes affect both slices.",
	Example:        "s := []int{0, 1, 2, 3, 4, 5}\nsub := s[1:4]   // [1, 2, 3]\nstart := s[:3]  // [0, 1, 2]\nend := s[2:]    // [2, 3, 4, 5]\nall := s[:]     // copy of s",
	UseCase:        "Use slicing to extract portions of slices, implement windowing/pagination, or work with subsets of data. Be aware that slices share underlying storage - modifying one affects the other until a reallocation occurs.",
	Prerequisites:  []string{"slice-basics"},
	RelatedTopics:  []string{"slice-basics", "slice-copy", "array-declaration"},
	DocsURL:        "https://go.dev/tour/moretypes/7",
}

func init() {
	Register(Concept017)
}
