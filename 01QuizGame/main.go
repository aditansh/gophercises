package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/aditansh/gophercises/01QuizGame/problems"
	"github.com/aditansh/gophercises/01QuizGame/quiz"
)

const (
	FileFlag      = "file"
	FileFlagValue = "problems.csv"
	FileFlagUsage = "a csv file in the format of 'question,answer'"

	timeFlag      = "time"
	timeFlagValue = 30
	timeFlagUsage = "the time limit for the quiz in seconds"
)

type Flagger interface {
	StringVar(p *string, name, value, usage string)
	IntVar(p *int, name string, value int, usage string)
}

type quizFlags struct{}

func (q *quizFlags) StringVar(p *string, name, value, usage string) {
	flag.StringVar(p, name, value, usage)
}

func (q *quizFlags) IntVar(p *int, name string, value int, usage string) {
	flag.IntVar(p, name, value, usage)
}

type Timer interface {
	NewTimer(d time.Duration) *time.Timer
}

type quizTimer struct{}

func (q quizTimer) NewTimer(d time.Duration) *time.Timer {
	return time.NewTimer(d)
}

func ReadCSV(r io.Reader) quiz.Quiz {
	csvReader := csv.NewReader(r)

	problem := []problems.Problem{}
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		problem = append(problem, problems.NewProblem(record))
	}

	return quiz.NewQuiz(problem)
}

var TimerSeconds int
var file string

func ConfigFlags(f Flagger) {
	f.StringVar(&file, FileFlag, FileFlagValue, FileFlagUsage)
	f.IntVar(&TimerSeconds, timeFlag, timeFlagValue, timeFlagUsage)
}

func StartTimer(w io.Writer, r io.Reader, timer Timer) *time.Timer {
	fmt.Fprintln(w, "Press any key to start the quiz")
	fmt.Fscanln(r)

	return timer.NewTimer(time.Duration(TimerSeconds) * time.Second)
}

func init() {
	flagger := &quizFlags{}
	ConfigFlags(flagger)

	flag.Parse()
}

func main() {
	file, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	quiz := ReadCSV(file)

	timer := StartTimer(os.Stdout, os.Stdin, quizTimer{})

	go func() {
		<-timer.C
		fmt.Println("\nTime's up!")
		quiz.GetScore(os.Stdout)
		os.Exit(0)
	}()

	quiz.Play(os.Stdout, os.Stdin)
}
