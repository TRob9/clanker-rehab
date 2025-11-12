package concepts

// 018. Copy Slice
var Concept018 = Concept{
	Number:      18,
	ID:          "slice-copy",
	Category:    "Data Structures",
	Name:        "18. Copy Slice",
	Description: "Use copy to duplicate slice",
	Instruction: "Create a source slice containing 1 and 2, create a destination slice with length 2 using make, copy the source to the destination, then print the second element of the destination",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	src := []int{1, 2}
	dst := make([]int, 2)
	copy(dst, src)
	fmt.Println(dst[1])
}`,
	ExpectedOutput: "2",
	Difficulty:     "beginner",
	Explanation:    "The copy() built-in function copies elements from source to destination slice. It returns the number of elements copied (the minimum of len(src) and len(dst)). copy() prevents shared underlying array issues.",
	Example:        "src := []int{1, 2, 3}\ndst := make([]int, len(src))\nn := copy(dst, src)  // n = 3\n// Partial copy:\npartial := make([]int, 2)\ncopy(partial, src)   // [1, 2]",
	UseCase:        "Use copy when you need an independent duplicate of a slice, to avoid shared underlying array problems, or when implementing algorithms that require immutable data. copy is safer than slicing when you need true independence.",
	Prerequisites:  []string{"slice-basics", "slice-make"},
	RelatedTopics:  []string{"slice-basics", "slice-slicing", "slice-make"},
	DocsURL:        "https://go.dev/pkg/builtin/#copy",
}

func init() {
	Register(Concept018)
}
