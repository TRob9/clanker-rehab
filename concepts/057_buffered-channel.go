package concepts

// 057. Buffered Channel
var Concept057 = Concept{
	Number:      57,
	ID:          "buffered-channel",
	Category:    "Concurrency",
	Name:        "57. Buffered Channel",
	Description: "Create channel with buffer",
	Instruction: "Create a buffered channel of integers with capacity 2, send 1 and 2 to the channel, then receive and print the first value",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
}`,
	ExpectedOutput: "1",
	Difficulty:     "beginner",
	Explanation:    "Buffered channels have capacity: make(chan T, capacity). Sends block only when buffer is full; receives block only when buffer is empty. This allows asynchronous send/receive up to the buffer limit.",
	Example:        "ch := make(chan int, 3)  // capacity 3\n\nch <- 1  // doesn't block\nch <- 2  // doesn't block\nch <- 3  // doesn't block\n// ch <- 4  // would block (buffer full)\n\nfmt.Println(<-ch)  // 1\nfmt.Println(<-ch)  // 2\n// Now buffer has space",
	UseCase:        "Use buffered channels to reduce blocking, implement producer-consumer patterns, create work queues, or handle bursts. Buffer size should match your throughput needs. Unbuffered (capacity 0) ensures rendezvous.",
	Prerequisites:  []string{"channel-create"},
	RelatedTopics:  []string{"channel-create", "channel-send-receive"},
	DocsURL:        "https://go.dev/tour/concurrency/3",
}

func init() {
	Register(Concept057)
}
