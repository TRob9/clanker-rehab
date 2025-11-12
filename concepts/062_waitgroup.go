package concepts

// 062. WaitGroup Basics
var Concept062 = Concept{
	Number:      62,
	ID:          "waitgroup",
	Category:    "Concurrency",
	Name:        "62. WaitGroup Basics",
	Description: "Use sync.WaitGroup to wait for goroutines",
	Instruction: "Declare a WaitGroup, add 1 to it, launch a goroutine that defers calling Done on the WaitGroup and prints the word "done", then call Wait on the WaitGroup in main",
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
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("done")
	}()
	wg.Wait()
}`,
	ExpectedOutput: "done",
	Difficulty:     "beginner",
	Explanation:    "WaitGroup waits for a collection of goroutines to finish. Add(n) increments the counter by n, Done() decrements it by 1, and Wait() blocks until the counter reaches zero.",
	Example:        "var wg sync.WaitGroup\n\nfor i := 0; i < 5; i++ {\n    wg.Add(1)\n    go func(id int) {\n        defer wg.Done()\n        fmt.Println(\"worker\", id)\n    }(i)\n}\n\nwg.Wait()  // blocks until all 5 call Done()\nfmt.Println(\"all done\")",
	UseCase:        "Use WaitGroup to wait for multiple goroutines to complete, synchronize parallel tasks, or ensure cleanup happens after all workers finish. Essential pattern for parallel processing and fan-out/fan-in.",
	Prerequisites:  []string{"goroutine", "defer"},
	RelatedTopics:  []string{"goroutine", "defer", "mutex"},
	DocsURL:        "https://pkg.go.dev/sync#WaitGroup",
}

func init() {
	Register(Concept062)
}
