package quiz

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/aditansh/gophercises/01QuizGame/problems"
)

func TestNewQuiz(t *testing.T) {
	problems := sampleProblems()

	want := Quiz{problems, 0}
	got := NewQuiz(problems)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Failed to create a new quiz. Got: %v, Want: %v", got, want)
	} else {
		t.Log("Successfully created a new quiz. Got: ", got, "Want: ", want)
	}
}

func TestPlay(t *testing.T) {
	t.Run("Stating quiz", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		quiz := sampleQuiz()
		runQuiz(buffer, &quiz)

		want := 2
		got := quiz.Score

		if got != want {
			t.Errorf("Expected correct %d, got correct %d", want, got)
		}

		expectedOutput := "1+1\nEnter your answer: 2+2\nEnter your answer: You got 2 questions right!\n"

		if buffer.String() != expectedOutput {
			t.Errorf("Expected output %q, got %q", expectedOutput, buffer.String())
		}
	})
}

func sampleProblems() []problems.Problem {

	record1 := []string{"1+1", "2"}
	record2 := []string{"2+2", "4"}

	return []problems.Problem{
		problems.NewProblem(record1),
		problems.NewProblem(record2),
	}
}

func sampleQuiz() Quiz {
	problems := sampleProblems()
	return NewQuiz(problems)
}

func runQuiz(buffer io.Writer, quiz *Quiz) {
	ans := strings.NewReader("2\n4\n")
	quiz.Play(buffer, ans)
}
