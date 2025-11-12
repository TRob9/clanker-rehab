package concepts

// 016. Append to Slice
var Concept016 = Concept{
	Number:      16,
	ID:          "slice-append",
	Category:    "Data Structures",
	Name:        "16. Append to Slice",
	Description: "Use append to add elements",
	Instruction: "Create a slice containing 1, append the value 2 to it, then print the length of the slice",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	s := []int{1}
	s = append(s, 2)
	fmt.Println(len(s))
}`,
	ExpectedOutput: "2",
	Difficulty:     "beginner",
	Explanation:    "The append() built-in function adds elements to a slice, returning a new slice. If the underlying array is full, append allocates a new, larger array and copies the elements. Always assign the result back: s = append(s, elem).",
	Example:        "s := []int{1, 2}\ns = append(s, 3)           // [1, 2, 3]\ns = append(s, 4, 5, 6)     // [1, 2, 3, 4, 5, 6]\n// Append slice to slice:\ns = append(s, []int{7,8}...)",
	UseCase:        "Use append to grow slices dynamically, build lists incrementally, or combine slices. append is efficient - it only allocates when necessary. Remember that append may return a new slice if reallocation occurs.",
	Prerequisites:  []string{"slice-basics"},
	RelatedTopics:  []string{"slice-basics", "slice-make", "slice-capacity"},
	DocsURL:        "https://go.dev/tour/moretypes/15",
}

func init() {
	Register(Concept016)
}
