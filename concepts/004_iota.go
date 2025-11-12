package concepts

// 004. Iota Enumerator
var Concept004 = Concept{
	Number:      4,
	ID:          "iota",
	Category:    "Core Syntax",
	Name:        "4. Iota Enumerator",
	Description: "Use iota for incrementing constants",
	Instruction: "Create three constants named A, B, and C using the iota enumerator. Print the value of the constant B",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	const (
		A = iota
		B
		C
	)
	fmt.Println(B)
}`,
	ExpectedOutput: "1",
	Difficulty:     "beginner",
	Explanation:    "iota is a special identifier used in constant declarations to create incrementing values. It starts at 0 and increments by 1 for each line in a const block. iota resets to 0 whenever the word 'const' appears.",
	Example:        "const (\n    A = iota  // 0\n    B         // 1\n    C         // 2\n    D         // 3\n)",
	UseCase:        "Use iota for creating enumerations and sequential constants like days of the week, status codes, or bit flags. It's especially powerful when combined with bit shifting for creating bitmask flags.",
	Prerequisites:  []string{"constants"},
	RelatedTopics:  []string{"constants", "const-block", "iota-bitmask"},
	DocsURL:        "https://go.dev/ref/spec#Iota",
}

func init() {
	Register(Concept004)
}
