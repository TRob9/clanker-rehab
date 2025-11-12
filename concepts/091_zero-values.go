package concepts

// 091. Zero Values
var Concept091 = Concept{
	Number:      91,
	ID:          "zero-values",
	Category:    "Miscellaneous",
	Name:        "91. Zero Values",
	Description: "Uninitialized variables have zero values",
	Instruction: "Declare an uninitialized integer variable x, string variable s, and boolean variable b, then print the result of checking if x equals 0, s equals empty string, and b is false",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	var x int
	var s string
	var b bool
	fmt.Println(x == 0 && s == "" && b == false)
}`,
	ExpectedOutput: "true",
	Difficulty:     "beginner",
	Explanation:    "In Go, variables declared without an explicit initial value are given their zero value. The zero value is type-specific: 0 for numeric types, false for booleans, \"\" for strings, and nil for pointers, slices, maps, channels, functions, and interfaces.",
	Example:        "var i int       // 0\nvar f float64   // 0.0\nvar b bool      // false\nvar s string    // \"\"\nvar p *int      // nil\nvar slice []int // nil\nvar m map[string]int // nil\n\n// Zero values make code simpler:\nvar counter int  // starts at 0, ready to use\ncounter++        // now 1",
	UseCase:        "Zero values eliminate the need for explicit initialization in many cases, making Go code cleaner and safer. Leverage zero values in your designs - for example, sync.Mutex's zero value is an unlocked mutex, ready to use.",
	Prerequisites:  []string{"var-declaration"},
	RelatedTopics:  []string{"var-declaration", "pointer-basics", "slice-basics"},
	DocsURL:        "https://go.dev/tour/basics/12",
}

func init() {
	Register(Concept091)
}
