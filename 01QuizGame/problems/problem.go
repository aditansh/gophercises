package problems

import (
	"fmt"
	"io"
)

type Problem struct {
	Question string
	Answer   string
	Options  []string
}

func NewProblem(record []string) Problem {
	return Problem{
		Question: record[0],
		Answer:   record[1],
		Options:  record[2:],
	}
}

func (p Problem) AskQuestion(w io.Writer) {
	w.Write([]byte(p.Question + "\n"))
	for _, option := range p.Options {
		w.Write([]byte(option + "\n"))
	}
	w.Write([]byte("Enter your answer: "))
}

func (p Problem) CheckAnswer(answer string) bool {
	return p.Answer == answer
}

func ReadAnswer(r io.Reader) (string, error) {
	var answer string
	_, err := fmt.Fscanln(r, &answer)
	return answer, err
}
