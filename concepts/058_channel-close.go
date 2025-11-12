package concepts

// 058. Close Channel
var Concept058 = Concept{
	Number:      58,
	ID:          "channel-close",
	Category:    "Concurrency",
	Name:        "58. Close Channel",
	Description: "Close channel with close()",
	Instruction: "Create a channel of integers, close it immediately, then use the comma-ok idiom to receive from it by ignoring the value and capturing ok, then print ok",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)
	_, ok := <-ch
	fmt.Println(ok)
}`,
	ExpectedOutput: "false",
	Difficulty:     "beginner",
	Explanation:    "close(ch) closes a channel - no more values can be sent. Receivers can test if a channel is closed using the comma-ok idiom: value, ok := <-ch. When closed, ok is false. Receiving from a closed channel returns the zero value.",
	Example:        "ch := make(chan int)\n\ngo func() {\n    ch <- 1\n    ch <- 2\n    close(ch)  // signal no more values\n}()\n\nfor v := range ch {  // loops until closed\n    fmt.Println(v)\n}\n\n// Or manually:\nv, ok := <-ch\nif !ok {\n    // channel closed\n}",
	UseCase:        "Use close to signal completion, end range loops over channels, or indicate no more data. Only senders should close channels. Closing is optional - channels are garbage collected. Common pattern: close after sending all values.",
	Prerequisites:  []string{"channel-create"},
	RelatedTopics:  []string{"range-channel", "channel-send-receive"},
	DocsURL:        "https://go.dev/tour/concurrency/4",
}

func init() {
	Register(Concept058)
}
