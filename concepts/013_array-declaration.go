package concepts

// 013. Array Declaration
var Concept013 = Concept{
	Number:      13,
	ID:          "array-declaration",
	Category:    "Data Structures",
	Name:        "13. Array Declaration",
	Description: "Fixed-size array",
	Instruction: "Declare an array of 3 integers containing the values 1, 2, and 3, then print the element at index 1",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr[1])
}`,
	ExpectedOutput: "2",
	Difficulty:     "beginner",
	Explanation:    "Arrays are fixed-length sequences of elements of a single type. The array's length is part of its type ([3]int is different from [4]int). Arrays are value types - they're copied when assigned or passed to functions.",
	Example:        "var arr [3]int = [3]int{1, 2, 3}\narr2 := [5]string{\"a\", \"b\", \"c\", \"d\", \"e\"}\n// Length inferred:\narr3 := [...]int{10, 20, 30}",
	UseCase:        "Arrays are rarely used directly in Go. Use slices instead for most cases. Arrays are useful when you need a fixed size known at compile time, or for optimization in performance-critical code.",
	Prerequisites:  []string{"var-declaration"},
	RelatedTopics:  []string{"slice-basics", "slice-make"},
	DocsURL:        "https://go.dev/tour/moretypes/6",
}

func init() {
	Register(Concept013)
}
