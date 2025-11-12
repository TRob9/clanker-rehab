package concepts

// 040. Pointer to Struct
var Concept040 = Concept{
	Number:      40,
	ID:          "pointer-struct",
	Category:    "Pointers & Methods",
	Name:        "40. Pointer to Struct",
	Description: "Access struct fields via pointer",
	Instruction: "Define type T as a struct with field X of type int, create an instance t with X set to 5, create a pointer p to t, use p to set X to 10, then print the X field of t",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	type T struct { X int }
	t := T{X: 5}
	p := &t
	p.X = 10
	fmt.Println(t.X)
}`,
	ExpectedOutput: "10",
	Difficulty:     "beginner",
	Explanation:    "Go provides automatic dereferencing for struct pointers. You can use p.X instead of (*p).X to access fields. This syntactic sugar makes pointer usage cleaner. Struct pointers are very common for avoiding large struct copies.",
	Example:        "type Person struct {\n    Name string\n    Age  int\n}\np := &Person{\"Alice\", 30}\np.Age = 31  // automatically dereferences\n// Same as (*p).Age = 31\n\nfunc birthday(p *Person) {\n    p.Age++  // modifies original\n}",
	UseCase:        "Use struct pointers when passing structs to functions to avoid copying, when you need to modify struct fields, and for all methods that mutate state. Most Go code uses pointers for structs larger than a few fields.",
	Prerequisites:  []string{"pointer-basics", "struct-declaration"},
	RelatedTopics:  []string{"method-pointer-receiver", "dereference"},
	DocsURL:        "https://go.dev/tour/moretypes/4",
}

func init() {
	Register(Concept040)
}
