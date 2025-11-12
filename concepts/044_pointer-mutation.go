package concepts

// 044. Pointer Receiver Mutation
var Concept044 = Concept{
	Number:      44,
	ID:          "pointer-mutation",
	Category:    "Pointers & Methods",
	Name:        "44. Pointer Receiver Mutation",
	Description: "Modify struct via pointer receiver",
	Instruction: "Define type Box as a struct with field W of type int, implement method Resize with parameter w (int) and pointer receiver that sets the W field to w, create a box with W set to 5, call the Resize method with 10, then print the W field",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

type Box struct { W int }

func (b *Box) Resize(w int) { b.W = w }

func main() {
	box := Box{W: 5}
	box.Resize(10)
	fmt.Println(box.W)
}`,
	ExpectedOutput: "10",
	Difficulty:     "beginner",
	Explanation:    "Pointer receiver methods can mutate the receiver's state. Changes made inside the method affect the original instance. This is the standard pattern for methods that update object state, like setters, incrementers, or state machines.",
	Example:        "type Account struct {\n    Balance int\n}\n\nfunc (a *Account) Deposit(amount int) {\n    a.Balance += amount  // mutates original\n}\n\nfunc (a *Account) Withdraw(amount int) bool {\n    if a.Balance >= amount {\n        a.Balance -= amount\n        return true\n    }\n    return false\n}\n\nacc := Account{Balance: 100}\nacc.Deposit(50)\nacc.Balance  // 150",
	UseCase:        "Use pointer receivers for methods that need to modify state, implement builders or fluent interfaces, or update configuration. This is the standard pattern for mutable objects in Go.",
	Prerequisites:  []string{"method-pointer-receiver"},
	RelatedTopics:  []string{"method-pointer-receiver", "dereference", "value-vs-reference"},
	DocsURL:        "https://go.dev/tour/methods/4",
}

func init() {
	Register(Concept044)
}
