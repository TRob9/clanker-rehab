package concepts

// 063. Mutex Basics
var Concept063 = Concept{
	Number:      63,
	ID:          "mutex",
	Category:    "Concurrency",
	Name:        "63. Mutex Basics",
	Description: "Use sync.Mutex for mutual exclusion",
	Instruction: "Declare a Mutex, initialize a variable to 0, lock the mutex, increment the variable, unlock the mutex, then print the variable",
	Boilerplate: `package main

import (
	"fmt"
	"sync"
)

func main() {
	// Your code here
}`,
	Answer: `package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	x := 0
	mu.Lock()
	x++
	mu.Unlock()
	fmt.Println(x)
}`,
	ExpectedOutput: "1",
	Difficulty:     "beginner",
	Explanation:    "Mutex provides mutual exclusion - only one goroutine can hold the lock at a time. Lock() acquires the lock (blocks if held), Unlock() releases it. Use to protect shared data from concurrent access.",
	Example:        "var mu sync.Mutex\nvar counter int\n\nfunc increment() {\n    mu.Lock()\n    defer mu.Unlock()  // always unlock\n    counter++\n}\n\n// Multiple goroutines safely share counter:\nfor i := 0; i < 100; i++ {\n    go increment()\n}",
	UseCase:        "Use mutexes to protect shared state, prevent data races, and ensure atomic operations on non-atomic data. Always use defer mu.Unlock() to prevent deadlocks. Prefer channels when possible - mutexes are for protecting invariants.",
	Prerequisites:  []string{"goroutine"},
	RelatedTopics:  []string{"waitgroup", "atomic-operations"},
	DocsURL:        "https://go.dev/tour/concurrency/9",
}

func init() {
	Register(Concept063)
}
