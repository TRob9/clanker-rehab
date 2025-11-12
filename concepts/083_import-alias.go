package concepts

// 083. Import Alias
var Concept083 = Concept{
	Number:      83,
	ID:          "import-alias",
	Category:    "Tooling & Tests",
	Name:        "83. Import Alias",
	Description: "Alias imported package",
	Instruction: "Import the fmt package with an alias f, then use the Println function from the aliased package to print "alias"",
	Boilerplate: `package main

func main() {
	// Your code here
}`,
	Answer: `package main

import f "fmt"

func main() {
	f.Println("alias")
}`,
	ExpectedOutput: "alias",
	Difficulty:     "beginner",
	Explanation:    "Import aliases let you rename packages on import. Use an identifier before the import path. Helpful for avoiding name conflicts, shortening long package names, or using multiple versions of the same package.",
	Example:        "import (\n    f \"fmt\"           // alias fmt as f\n    crand \"crypto/rand\"\n    mrand \"math/rand\"\n)\n\nf.Println(\"using alias\")\ncrand.Read(...)  // crypto/rand\nmrand.Intn(10)   // math/rand",
	UseCase:        "Use import aliases to resolve package name conflicts (crypto/rand vs math/rand), shorten long package names for readability, or when package name doesn't match its purpose. Keep aliases intuitive - others need to understand your code.",
	Prerequisites:  nil,
	RelatedTopics:  []string{"blank-import"},
	DocsURL:        "https://go.dev/ref/spec#Import_declarations",
}

func init() {
	Register(Concept083)
}
