package concepts

// 027. Range over Slice
var Concept027 = Concept{
	Number:      27,
	ID:          "range-slice",
	Category:    "Data Structures",
	Name:        "27. Range over Slice",
	Description: "Iterate with range",
	Instruction: "Create a slice containing 10 and 20, use a for-range loop to iterate over it, and print each value on a separate line",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	s := []int{10, 20}
	for _, v := range s {
		fmt.Println(v)
	}
}`,
	ExpectedOutput: "10\n20",
	Difficulty:     "beginner",
	Explanation:    "The range keyword iterates over slices, arrays, maps, and channels. For slices/arrays, range returns index and value. Use the blank identifier _ to ignore unwanted values.",
	Example:        "nums := []int{10, 20, 30}\nfor i, v := range nums {\n    fmt.Printf(\"%d: %d\\\n\", i, v)\n}\n// Ignore index:\nfor _, v := range nums {\n    fmt.Println(v)\n}\n// Index only:\nfor i := range nums {\n    fmt.Println(i)\n}",
	UseCase:        "Use range for clean iteration over collections. It's more idiomatic than index-based loops and handles length automatically. range works with slices, arrays, strings, maps, and channels.",
	Prerequisites:  []string{"slice-basics", "for-loop"},
	RelatedTopics:  []string{"for-loop", "slice-basics", "blank-identifier"},
	DocsURL:        "https://go.dev/tour/moretypes/16",
}

func init() {
	Register(Concept027)
}
