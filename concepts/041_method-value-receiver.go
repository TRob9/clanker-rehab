package concepts

// 041. Method with Value Receiver
var Concept041 = Concept{
	Number:      41,
	ID:          "method-value-receiver",
	Category:    "Pointers & Methods",
	Name:        "41. Method with Value Receiver",
	Description: "Define method on type",
	Instruction: "Implement the Double method on type Num that returns the value multiplied by 2",
	Boilerplate: `package main

import "fmt"

type Num int

func (n Num) Double() int {
	// Your code here
}

func main() {
	x := Num(3)
	fmt.Println(x.Double())
	y := Num(10)
	fmt.Println(y.Double())
}`,
	Answer: `package main

import "fmt"

type Num int

func (n Num) Double() int {
	return int(n) * 2
}

func main() {
	x := Num(3)
	fmt.Println(x.Double())
	y := Num(10)
	fmt.Println(y.Double())
}`,
	ExpectedOutput: "6\n20",
	TestCases: []TestCase{
		{Input: "Num(3).Double()", Expected: "6"},
		{Input: "Num(10).Double()", Expected: "20"},
	},
	Difficulty:     "beginner",
	Explanation:    "Methods are functions with a receiver argument. Value receivers get a copy of the value. Methods with value receivers cannot modify the original value - they work on a copy. Use for read-only operations and small types.",
	Example:        "type Point struct { X, Y int }\n\nfunc (p Point) Distance() float64 {\n    return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))\n}\n\np := Point{3, 4}\ndist := p.Distance()  // 5.0\n// p is copied, original unchanged",
	UseCase:        "Use value receivers for small types (primitives, small structs), when the method doesn't need to modify the receiver, or when you want to work with copies. Value receivers are safer - no nil pointer issues.",
	Prerequisites:  []string{"struct-declaration"},
	RelatedTopics:  []string{"method-pointer-receiver", "pointer-struct"},
	DocsURL:        "https://go.dev/tour/methods/4",
}

func init() {
	Register(Concept041)
}
