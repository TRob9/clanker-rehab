package concepts

// 084. Blank Import
var Concept084 = Concept{
	Number:      84,
	ID:          "blank-import",
	Category:    "Tooling & Tests",
	Name:        "84. Blank Import",
	Description: "Import for side effects with _",
	Instruction: "Import fmt using the blank identifier for side effects only, then print "ok" using the Println function",
	Boilerplate: `package main

import (
	_ "crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println("ok")
}`,
	Answer: `package main

import (
	_ "crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println("ok")
}`,
	ExpectedOutput: "ok",
	Difficulty:     "beginner",
	Explanation:    "Blank imports (import _ \"package\") import a package only for its side effects (init functions), without using its exported identifiers. The package's init() functions run, but you can't reference the package directly.",
	Example:        "import (\n    _ \"github.com/lib/pq\"  // Register PostgreSQL driver\n    \"database/sql\"\n)\n\n// pq.init() ran, registering \"postgres\" driver\ndb, err := sql.Open(\"postgres\", connStr)\n\n// Also used for image formats:\nimport _ \"image/png\"  // Registers PNG decoder",
	UseCase:        "Use blank imports to register drivers (database, image formats), load plugins, or run package init functions. Common with database/sql drivers and image decoders. The import ensures the package's init() runs but you don't directly use its code.",
	Prerequisites:  []string{"blank-identifier"},
	RelatedTopics:  []string{"init-function", "blank-identifier"},
	DocsURL:        "https://go.dev/doc/effective_go#blank_import",
}

func init() {
	Register(Concept084)
}
