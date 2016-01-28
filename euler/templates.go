package main

var (
	problemTemplate = `/*
{{problem_description}}
*/
package main

import "github.com/jacobhands/pu"

var (
	answer  = "NA" // Change to correct answer once solved.
	problem = pu.Problem{ID: {{problem_id}}, Solver: solveProblem{{problem_id}}, CorrectAnswer: answer}
)

func main() {
	problem.Answer()
}

// This should return the answer in the same formatting as 'answer' is set to above.
func solveProblem{{problem_id}}() string {
	return ""
}
`
	problemTestTemplate = `package main

import "testing"

func BenchmarkSolveProblem{{problem_id}}(b *testing.B) {
	problem.Bench(b)
}

func TestSolveProblem{{problem_id}}(t *testing.T) {
	problem.Test(t)
}
`
	problemIDString          = "{{problem_id}}"
	problemDescriptionString = "{{problem_description}}"
)
