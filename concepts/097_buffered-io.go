package concepts

// 097. Buffered I/O with bufio
var Concept097 = Concept{
	Number:      97,
	ID:          "buffered-io",
	Category:    "Standard Library",
	Name:        "97. Buffered I/O with bufio",
	Description: "Use bufio for efficient buffered reading",
	Instruction: "Create a strings Reader with the word test, wrap it in a buffered reader using NewReader from the bufio package, read a line with the ReadString method using newline as delimiter, and print the trimmed result.",
	Boilerplate: `package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	// Your code here
}`,
	Answer: `package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("test"))
	line, _ := reader.ReadString('\n')
	fmt.Println(strings.TrimSpace(line))
}`,
	ExpectedOutput: "test",
	Difficulty:     "advanced",
	Explanation:    "The bufio package provides buffered I/O operations for more efficient reading and writing. bufio.NewReader wraps an io.Reader and adds buffering, reducing the number of system calls. Methods like ReadString read until a delimiter.",
	Example:        "file, _ := os.Open(\"file.txt\")\ndefer file.Close()\nreader := bufio.NewReader(file)\n\n// Read line by line:\nfor {\n    line, err := reader.ReadString('\\\n')\n    if err != nil { break }\n    fmt.Print(line)\n}\n\n// Or use Scanner:\nscanner := bufio.NewScanner(file)\nfor scanner.Scan() {\n    fmt.Println(scanner.Text())\n}",
	UseCase:        "Use bufio for reading files line-by-line, parsing large files efficiently, or reducing I/O overhead. bufio.Scanner is easier for simple line reading. bufio.Writer provides buffered writing for better performance when making many small writes.",
	Prerequisites:  nil,
	RelatedTopics:  nil,
	DocsURL:        "https://pkg.go.dev/bufio",
}

func init() {
	Register(Concept097)
}
