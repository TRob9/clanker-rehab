package concepts

// 104. Slice Capacity vs Length
var Concept104 = Concept{
	Number:      104,
	ID:          "slice-capacity",
	Category:    "Data Structures",
	Name:        "104. Slice Capacity vs Length",
	Description: "Understand difference between len and cap",
	Instruction: "Create a slice using make with length 2 and capacity 5. Print the capacity of the slice.",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	s := make([]int, 2, 5)
	fmt.Println(cap(s))
}`,
	ExpectedOutput: "5",
	Difficulty:     "advanced",
	Explanation:    "Slices have both length (len) and capacity (cap). Length is the number of elements it contains. Capacity is the number of elements in the underlying array from the slice's first element. Capacity determines when reallocation occurs.",
	Example:        "s := make([]int, 3, 5)\nfmt.Println(len(s))  // 3\nfmt.Println(cap(s))  // 5\ns = append(s, 1, 2)  // len=5, cap=5, no realloc\ns = append(s, 3)     // len=6, cap=10, realloc!",
	UseCase:        "Understanding capacity helps optimize memory and performance. Preallocate capacity when size is known to avoid multiple reallocations during append. Check capacity before append in performance-critical code.",
	Prerequisites:  []string{"slice-basics", "slice-make"},
	RelatedTopics:  []string{"slice-make", "slice-append"},
	DocsURL:        "https://go.dev/blog/slices-intro",
}

func init() {
	Register(Concept104)
}
