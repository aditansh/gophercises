package quiz

import (
	"fmt"
	"io"

	"github.com/aditansh/gophercises/01QuizGame/problems"
)

type Quiz struct {
	Problems []problems.Problem
	Score    int
}

func NewQuiz(problems []problems.Problem) Quiz {
	return Quiz{
		Problems: problems,
		Score:    0,
	}
}

func (q *Quiz) Play(w io.Writer, r io.Reader) {
	for _, problem := range q.Problems {
		problem.AskQuestion(w)
		answer, err := problems.ReadAnswer(r)
		if err != nil {
			panic(err)
		}
		if problem.CheckAnswer(answer) {
			q.Score++
		}
	}

	q.GetScore(w)
}

func (q Quiz) GetScore(w io.Writer) {
	fmt.Fprintf(w, "You got %d questions right!\n", q.Score)
}
