package concepts

// 059. Range over Channel
var Concept059 = Concept{
	Number:      59,
	ID:          "range-channel",
	Category:    "Concurrency",
	Name:        "59. Range over Channel",
	Description: "Iterate channel with range",
	Instruction: "Create a buffered channel of integers with capacity 2, send 1 and 2 to it, close the channel, then use a for-range loop to iterate over the channel and print each value",
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
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
}`,
	ExpectedOutput: "1\n2",
	Difficulty:     "beginner",
	Explanation:    "Range loops over channels, receiving values until the channel is closed. The loop automatically terminates when the channel is closed. This is cleaner than manually checking with comma-ok idiom.",
	Example:        "ch := make(chan int)\n\ngo func() {\n    for i := 0; i < 5; i++ {\n        ch <- i\n    }\n    close(ch)  // must close to end range\n}()\n\nfor val := range ch {\n    fmt.Println(val)  // 0, 1, 2, 3, 4\n}\n// Loop ends when channel closes",
	UseCase:        "Use range over channels to consume all values from a channel until it's closed. Perfect for processing streams, consuming work queues, or collecting results. The channel must be closed to end the loop.",
	Prerequisites:  []string{"channel-close", "range-slice"},
	RelatedTopics:  []string{"channel-close", "channel-send-receive"},
	DocsURL:        "https://go.dev/tour/concurrency/4",
}

func init() {
	Register(Concept059)
}
