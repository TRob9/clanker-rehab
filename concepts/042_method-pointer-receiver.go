package concepts

// 042. Method with Pointer Receiver
var Concept042 = Concept{
	Number:      42,
	ID:          "method-pointer-receiver",
	Category:    "Pointers & Methods",
	Name:        "42. Method with Pointer Receiver",
	Description: "Mutate receiver with pointer",
	Instruction: "Implement the Inc method on Counter that increments the Val field by 1",
	Boilerplate: `package main

import "fmt"

type Counter struct {
	Val int
}

func (c *Counter) Inc() {
	// Your code here
}

func main() {
	x := Counter{}
	x.Inc()
	fmt.Println(x.Val)
	x.Inc()
	fmt.Println(x.Val)
}`,
	Answer: `package main

import "fmt"

type Counter struct {
	Val int
}

func (c *Counter) Inc() {
	c.Val++
}

func main() {
	x := Counter{}
	x.Inc()
	fmt.Println(x.Val)
	x.Inc()
	fmt.Println(x.Val)
}`,
	ExpectedOutput: "1\n2",
	TestCases: []TestCase{
		{Input: "Inc() once", Expected: "Val = 1"},
		{Input: "Inc() twice", Expected: "Val = 2"},
	},
	Difficulty:     "beginner",
	Explanation:    "Methods with pointer receivers can modify the original value. The receiver is a pointer, so changes affect the actual instance. Go automatically takes the address when calling pointer methods on values. Use pointer receivers for mutation.",
	Example:        "type Counter struct { Val int }\n\nfunc (c *Counter) Inc() {\n    c.Val++  // modifies original\n}\n\nfunc (c *Counter) Reset() {\n    c.Val = 0\n}\n\nc := Counter{}\nc.Inc()  // Go: (&c).Inc()\nc.Val  // 1",
	UseCase:        "Use pointer receivers when methods need to modify the receiver, for large structs (avoid copying), or for consistency (if one method uses pointer receiver, all should). Most struct methods use pointer receivers.",
	Prerequisites:  []string{"pointer-basics", "struct-declaration"},
	RelatedTopics:  []string{"method-value-receiver", "pointer-mutation", "pointer-struct"},
	DocsURL:        "https://go.dev/tour/methods/4",
}

func init() {
	Register(Concept042)
}
