package concepts

// 015. Slice with Make
var Concept015 = Concept{
	Number:      15,
	ID:          "slice-make",
	Category:    "Data Structures",
	Name:        "15. Slice with Make",
	Description: "Use make to create slice",
	Instruction: "Use the make function to create a slice of integers with length 2, set the first element to 5, then print the first element",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	s := make([]int, 2)
	s[0] = 5
	fmt.Println(s[0])
}`,
	ExpectedOutput: "5",
	Difficulty:     "beginner",
	Explanation:    "The make() function creates slices with a specified length and optional capacity. make([]T, length, capacity) allocates an underlying array and returns a slice pointing to it. If capacity is omitted, it equals length.",
	Example:        "s1 := make([]int, 5)        // len=5, cap=5\ns2 := make([]int, 3, 10)    // len=3, cap=10\ns3 := make([]string, 0, 5)  // empty but has capacity",
	UseCase:        "Use make when you know the size upfront (avoids reallocations), when you need specific capacity for performance, or when initializing a slice with zero values. Preallocating capacity can significantly improve append performance.",
	Prerequisites:  []string{"slice-basics"},
	RelatedTopics:  []string{"slice-basics", "slice-append", "slice-capacity", "new-vs-make"},
	DocsURL:        "https://go.dev/tour/moretypes/13",
}

func init() {
	Register(Concept015)
}
