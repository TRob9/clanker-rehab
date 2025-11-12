package concepts

// 075. sort.Ints
var Concept075 = Concept{
	Number:      75,
	ID:          "sort-ints",
	Category:    "Standard Library",
	Name:        "75. sort.Ints",
	Description: "Sort integer slice",
	Instruction: "Create s as a slice containing 3, 1, 2, use sort.Ints to sort it in-place, then print the first element",
	Boilerplate: `package main

import (
	"fmt"
	"sort"
)

func main() {
	// Your code here
}`,
	Answer: `package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{3, 1, 2}
	sort.Ints(s)
	fmt.Println(s[0])
}`,
	ExpectedOutput: "1",
	Difficulty:     "beginner",
	Explanation:    "sort.Ints sorts a slice of ints in ascending order, modifying the slice in-place. The sort package also provides sort.Float64s, sort.Strings, and the generic sort.Slice for custom sorting.",
	Example:        "nums := []int{5, 2, 8, 1, 9}\nsort.Ints(nums)  // nums is now [1, 2, 5, 8, 9]\n\n// For descending:\nsort.Sort(sort.Reverse(sort.IntSlice(nums)))\n\n// Check if sorted:\nsort.IntsAreSorted(nums)  // true",
	UseCase:        "Use sort.Ints for quick in-place sorting of integer slices. For custom comparisons or sorting structs, use sort.Slice with a custom less function. The sort is not stable (equal elements may be reordered).",
	Prerequisites:  []string{"slice-basics"},
	RelatedTopics:  []string{"custom-sort"},
	DocsURL:        "https://pkg.go.dev/sort#Ints",
}

func init() {
	Register(Concept075)
}
