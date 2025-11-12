package concepts

// 103. Iota with Bit Shifting
var Concept103 = Concept{
	Number:      103,
	ID:          "iota-bitmask",
	Category:    "Core Syntax",
	Name:        "103. Iota with Bit Shifting",
	Description: "Use iota with << for bitmask flags",
	Instruction: "Create constants using iota with bit shift operator: Read equals 1 left shift iota resulting in 1, Write resulting in 2, Execute resulting in 4. Print Write.",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	const (
		Read = 1 << iota
		Write
		Execute
	)
	fmt.Println(Write)
}`,
	ExpectedOutput: "2",
	Difficulty:     "advanced",
	Explanation:    "Combining iota with bit shifting (<<) creates power-of-two constants, perfect for bitmask flags. Each constant gets a unique bit: 1<<0=1, 1<<1=2, 1<<2=4, etc. This enables efficient flag combinations using bitwise operations.",
	Example:        "const (\n    Read = 1 << iota  // 1 (bit 0)\n    Write             // 2 (bit 1)\n    Execute           // 4 (bit 2)\n)\n\n// Combine flags with OR:\nperm := Read | Write  // 3 (bits 0 and 1)\n\n// Check flags with AND:\nif perm & Read != 0 {\n    fmt.Println(\"can read\")\n}\nif perm & Execute == 0 {\n    fmt.Println(\"cannot execute\")\n}",
	UseCase:        "Use iota bit shifting for file permissions, feature flags, option sets, or any scenario requiring multiple independent boolean flags packed efficiently. Common in system programming, configuration, and APIs that need compact flag representation.",
	Prerequisites:  []string{"iota", "constants"},
	RelatedTopics:  []string{"iota", "const-block"},
	DocsURL:        "https://go.dev/ref/spec#Iota",
}

func init() {
	Register(Concept103)
}
