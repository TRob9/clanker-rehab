package concepts

// 096. Custom Sort with sort.Slice
var Concept096 = Concept{
	Number:      96,
	ID:          "custom-sort",
	Category:    "Standard Library",
	Name:        "96. Custom Sort with sort.Slice",
	Description: "Sort a slice with custom comparison logic",
	Instruction: "Create an integer slice with values 3, 1, and 2, sort it in descending order using the Slice function from the sort package with a custom comparator, then print the first element (which should be 3).",
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
	nums := []int{3, 1, 2}
	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })
	fmt.Println(nums[0])
}`,
	ExpectedOutput: "3",
	Difficulty:     "advanced",
	Explanation:    "sort.Slice sorts any slice using a custom comparison function (less func). The less function receives two indices i, j and returns true if element i should come before j. This allows sorting by any criteria.",
	Example:        "nums := []int{3, 1, 4, 1, 5}\n// Descending:\nsort.Slice(nums, func(i, j int) bool {\n    return nums[i] > nums[j]  // > for descending\n})\n\n// Sort structs:\ntype Person struct { Name string; Age int }\npeople := []Person{{\"Alice\", 30}, {\"Bob\", 25}}\nsort.Slice(people, func(i, j int) bool {\n    return people[i].Age < people[j].Age  // sort by age\n})",
	UseCase:        "Use sort.Slice for custom sorting logic, sorting structs by specific fields, multi-key sorting, or when you need control over comparison. More flexible than sort.Ints/Strings. For stable sort (preserves order of equal elements), use sort.SliceStable.",
	Prerequisites:  []string{"slice-basics", "function-value"},
	RelatedTopics:  []string{"sort-ints"},
	DocsURL:        "https://pkg.go.dev/sort#Slice",
}

func init() {
	Register(Concept096)
}
