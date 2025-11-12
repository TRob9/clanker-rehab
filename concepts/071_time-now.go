package concepts

// 071. time.Now
var Concept071 = Concept{
	Number:      71,
	ID:          "time-now",
	Category:    "Standard Library",
	Name:        "71. time.Now",
	Description: "Get current time",
	Instruction: "Use the Now function from the time package to get the current time, then print whether the year is greater than 2000",
	Boilerplate: `package main

import (
	"fmt"
	"time"
)

func main() {
	// Your code here
}`,
	Answer: `package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Year() > 2000)
}`,
	ExpectedOutput: "true",
	Difficulty:     "beginner",
	Explanation:    "time.Now() returns the current local time as a time.Time value. The Time type has methods for extracting components (Year, Month, Day, Hour, etc.), formatting, and arithmetic.",
	Example:        "now := time.Now()\nnow.Year()              // 2024\nnow.Month()             // time.January (1)\nnow.Day()               // 15\nnow.Format(\"2006-01-02\") // \"2024-01-15\"\n// Time arithmetic:\ntomorrow := now.Add(24 * time.Hour)",
	UseCase:        "Use time.Now() for timestamps, logging, measuring elapsed time, scheduling, or any time-based logic. For UTC time, use time.Now().UTC(). For timestamps, use now.Unix() (seconds) or now.UnixNano() (nanoseconds).",
	Prerequisites:  nil,
	RelatedTopics:  nil,
	DocsURL:        "https://pkg.go.dev/time#Now",
}

func init() {
	Register(Concept071)
}
