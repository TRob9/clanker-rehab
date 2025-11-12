package concepts

// 100. Atomic Operations
var Concept100 = Concept{
	Number:      100,
	ID:          "atomic-operations",
	Category:    "Concurrency",
	Name:        "100. Atomic Operations",
	Description: "Use sync/atomic for lock-free operations",
	Instruction: "Declare a variable counter of type int64. Use the AddInt64 function from the atomic package to add 5 to it, then use the LoadInt64 function to read and print the value.",
	Boilerplate: `package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var counter int64
	// Your code here
}`,
	Answer: `package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var counter int64
	atomic.AddInt64(&counter, 5)
	fmt.Println(atomic.LoadInt64(&counter))
}`,
	ExpectedOutput: "5",
	Difficulty:     "advanced",
	Explanation:    "The sync/atomic package provides low-level atomic memory primitives for lock-free concurrent programming. Atomic operations are guaranteed to complete without interruption - no other goroutine can observe a half-completed operation. They're faster than mutexes for simple operations on single values.",
	Example:        "var counter int64\n\n// Multiple goroutines can safely increment:\nfor i := 0; i < 100; i++ {\n    go func() {\n        atomic.AddInt64(&counter, 1)\n    }()\n}\n\n// Safe read:\nval := atomic.LoadInt64(&counter)\n\n// Also: StoreInt64, SwapInt64, CompareAndSwapInt64",
	UseCase:        "Use atomic operations for simple shared counters, flags, or single-value state that multiple goroutines access. Perfect for metrics, rate limiters, and simple synchronization. Faster than mutexes but limited to primitive types. For complex state or multiple values, use mutexes or channels instead.",
	Prerequisites:  []string{"goroutine", "mutex"},
	RelatedTopics:  []string{"mutex", "waitgroup"},
	DocsURL:        "https://pkg.go.dev/sync/atomic",
}

func init() {
	Register(Concept100)
}
