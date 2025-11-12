package concepts

// 043. new vs make
var Concept043 = Concept{
	Number:      43,
	ID:          "new-vs-make",
	Category:    "Pointers & Methods",
	Name:        "43. new vs make",
	Description: "Understand new (returns pointer to zero value) vs make (initializes slices/maps/channels)",
	Instruction: "Use the new function to create a pointer to an integer, set the value at the pointer to 42 using dereference, then print the dereferenced value",
	Boilerplate: `package main

import "fmt"

func main() {
	// Your code here
}`,
	Answer: `package main

import "fmt"

func main() {
	p := new(int)
	*p = 42
	fmt.Println(*p)
}`,
	ExpectedOutput: "42",
	Difficulty:     "beginner",
	Explanation:    "new(T) allocates zeroed storage for a new item of type T and returns its address (*T). make(T, args) creates slices, maps, and channels only, returning an initialized (not zeroed) value of type T (not *T). new returns a pointer; make returns a value.",
	Example:        "// new - works for any type\np := new(int)       // p is *int, *p is 0\ns := new([]int)     // s is *[]int, *s is nil\n\n// make - only for slice/map/chan\nsl := make([]int, 5)     // slice, not pointer\nm := make(map[string]int) // map, not pointer",
	UseCase:        "Use new rarely - mostly for getting pointers to primitives. Use make for slices, maps, and channels (required). Prefer composite literals (&Type{}) over new for structs. new is less common than make in idiomatic Go.",
	Prerequisites:  []string{"pointer-basics", "slice-make"},
	RelatedTopics:  []string{"slice-make", "map-make", "pointer-basics"},
	DocsURL:        "https://go.dev/doc/effective_go#allocation_new",
}

func init() {
	Register(Concept043)
}
