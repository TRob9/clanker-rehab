package concepts

type TestCase struct {
	Input    string `json:"input"`
	Expected string `json:"expected"`
}

type Concept struct {
	Number         int        `json:"number"`         // LeetCode-style number
	ID             string     `json:"id"`
	Category       string     `json:"category"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	Instruction    string     `json:"instruction"`
	Boilerplate    string     `json:"boilerplate"`
	Answer         string     `json:"answer"`
	ExpectedOutput string     `json:"expectedOutput"`
	TestCases      []TestCase `json:"testCases,omitempty"`
	Difficulty     string     `json:"difficulty"`
	Explanation    string     `json:"explanation"`
	Example        string     `json:"example"`
	UseCase        string     `json:"useCase"`
	Prerequisites  []string   `json:"prerequisites"`
	RelatedTopics  []string   `json:"relatedTopics"`
	DocsURL        string     `json:"docsUrl"`
}

// Initialize allConcepts at package level so it's ready before any init() functions run
var allConcepts = make([]Concept, 0, 104)

func Register(c Concept) {
	allConcepts = append(allConcepts, c)
}

// GetAll returns all registered concepts
func GetAll() []Concept {
	return allConcepts
}
