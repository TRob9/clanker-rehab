package concepts

// 101. Method Expressions
var Concept101 = Concept{
	Number:      101,
	ID:          "method-expressions",
	Category:    "Functions & Closures",
	Name:        "101. Method Expressions",
	Description: "Call methods via method expressions (detached from receiver)",
	Instruction: "Implement the Inc method for the Counter type which should increment the Val field, then use the method expression syntax to call it and print the final Val field",
	Boilerplate: `package main

import "fmt"

type Counter struct {
	Val int
}

func (c *Counter) Inc() {
	// Your code here
}

func main() {
	c := &Counter{}
	fn := (*Counter).Inc
	fn(c)
	fmt.Println(c.Val)
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
	c := &Counter{}
	fn := (*Counter).Inc
	fn(c)
	fmt.Println(c.Val)
}`,
	ExpectedOutput: "1",
	TestCases: []TestCase{
		{Input: "Inc() once", Expected: "Val = 1"},
		{Input: "Inc() twice", Expected: "Val = 2"},
	},
	Difficulty:     "advanced",
	Explanation:    "Method expressions are methods accessed as functions, detached from their receiver. The syntax T.Method creates a function that takes the receiver as its first argument. This allows passing methods as function values.",
	Example:        "type Counter struct { Val int }\n\nfunc (c *Counter) Inc() { c.Val++ }\n\n// Method expression:\nfn := (*Counter).Inc  // func(*Counter)\nc := &Counter{}\nfn(c)  // calls c.Inc()\n\n// Useful for callbacks:\nfuncs := []func(*Counter){(*Counter).Inc, (*Counter).Reset}",
	UseCase:        "Use method expressions when you need to pass methods as function values, implement callbacks with methods, or create function tables. Particularly useful in reflection, generic code, and when building reusable function pipelines.",
	Prerequisites:  []string{"method-pointer-receiver", "function-value"},
	RelatedTopics:  []string{"function-value", "method-pointer-receiver"},
	DocsURL:        "https://go.dev/ref/spec#Method_expressions",
}

func init() {
	Register(Concept101)
}
