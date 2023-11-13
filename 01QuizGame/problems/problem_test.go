package problems

import (
	"reflect"
	"strings"
	"testing"
)

func TestNewProblem(t *testing.T) {
	record := []string{"question", "answer", "option1", "option2", "option3"}

	want := Problem{"question", "answer", []string{"option1", "option2", "option3"}}
	got := NewProblem(record)

	if got.Question != want.Question || got.Answer != want.Answer || !reflect.DeepEqual(got.Options, want.Options) {
		t.Errorf("Failed to create a new problem. Got: %v, Want: %v", got, want)
	} else {
		t.Log("Successfully created a new problem. Got: ", got, "Want: ", want)
	}
}

func TestAskQuestion(t *testing.T) {
	p := Problem{"question", "answer", []string{"option1", "option2", "option3"}}

	want := "question\noption1\noption2\noption3\nEnter your answer: "
	var got strings.Builder
	p.AskQuestion(&got)

	if got.String() != want {
		t.Errorf("Failed to ask question. Got: %v, Want: %v", got, want)
	} else {
		t.Log("Successfully asked question. Got: ", got, "Want: ", want)
	}
}

func TestCheckAnswer(t *testing.T) {
	p := Problem{"question", "answer", []string{"option1", "option2", "option3"}}

	want := true
	got := p.CheckAnswer("answer")

	if got != want {
		t.Errorf("Failed to check answer. Got: %v, Want: %v", got, want)
	} else {
		t.Log("Successfully checked answer. Got: ", got, "Want: ", want)
	}
}

func TestReadAnswer(t *testing.T) {
	want := "answer"
	got, err := ReadAnswer(strings.NewReader(want))

	if got != want || err != nil {
		t.Errorf("Failed to read answer. Got: %v, Want: %v", got, want)
	} else {
		t.Log("Successfully read answer. Got: ", got, "Want: ", want)
	}
}
