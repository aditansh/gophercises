package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"

	"github.com/aditansh/gophercises/01QuizGame/problems"
	"github.com/aditansh/gophercises/01QuizGame/quiz"
)

type flaggerMock struct {
	stringVarCalls  int
	intVarCalls     int
	varNames        []string
	varUsages       []string
	varStringValues []string
	varIntValues    []int
}

func (f *flaggerMock) StringVar(p *string, name, value, usage string) {
	f.stringVarCalls++
	f.varNames = append(f.varNames, name)
	f.varUsages = append(f.varUsages, usage)
	f.varStringValues = append(f.varStringValues, value)
}

func (f *flaggerMock) IntVar(p *int, name string, value int, usage string) {
	f.intVarCalls++
	f.varNames = append(f.varNames, name)
	f.varUsages = append(f.varUsages, usage)
	f.varIntValues = append(f.varIntValues, value)
}

type timerMock struct {
	duration int
}

func (t *timerMock) NewTimer(d time.Duration) *time.Timer {
	t.duration = int(d.Seconds())
	return time.NewTimer(1 * time.Millisecond)
}

func TestReadCSV(t *testing.T) {
	input := "1+1,2\n2+2,4\n"
	reader := bytes.NewReader([]byte(input))

	record1 := []string{"1+1", "2"}
	record2 := []string{"2+2", "4"}
	problem := []problems.Problem{
		problems.NewProblem(record1),
		problems.NewProblem(record2),
	}

	want := quiz.NewQuiz(problem)
	got := ReadCSV(reader)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Error reading CSV. Want: %v, Got: %v", want, got)
	} else {
		t.Log("ReadCSV test passed!")
	}
}

func TestConfigFlags(t *testing.T) {
	flagger := &flaggerMock{}
	ConfigFlags(flagger)

	wantStringVarCalls := 1
	gotStringVarCalls := flagger.stringVarCalls
	if wantStringVarCalls != gotStringVarCalls {
		t.Errorf("Error in StringVar ConfigFlags. Want: %v, Got: %v", wantStringVarCalls, gotStringVarCalls)
	}

	wantIntVarCalls := 1
	gotIntVarCalls := flagger.intVarCalls
	if wantIntVarCalls != gotIntVarCalls {
		t.Errorf("Error in IntVar ConfigFlags. Want: %v, Got: %v", wantIntVarCalls, gotIntVarCalls)
	}

	wantVarNames := []string{FileFlag, timeFlag}
	gotVarNames := flagger.varNames

	if !reflect.DeepEqual(wantVarNames, gotVarNames) {
		t.Errorf("Error in varNames ConfigFlags. Want: %v, Got: %v", wantVarNames, gotVarNames)
	}

	wantVarUsages := []string{FileFlagUsage, timeFlagUsage}
	gotVarUsages := flagger.varUsages

	if !reflect.DeepEqual(wantVarUsages, gotVarUsages) {
		t.Errorf("Error in varUsages ConfigFlags. Want: %v, Got: %v", wantVarUsages, gotVarUsages)
	}

	wantVarStringValues := []string{FileFlagValue}
	gotVarStringValues := flagger.varStringValues

	if !reflect.DeepEqual(wantVarStringValues, gotVarStringValues) {
		t.Errorf("Error in varStringValues ConfigFlags. Want: %v, Got: %v", wantVarStringValues, gotVarStringValues)
	}

	wantVarIntValues := []int{timeFlagValue}
	gotVarIntValues := flagger.varIntValues

	if !reflect.DeepEqual(wantVarIntValues, gotVarIntValues) {
		t.Errorf("Error in varIntValues ConfigFlags. Want: %v, Got: %v", wantVarIntValues, gotVarIntValues)
	}
}

func TestStartTimer(t *testing.T) {
	timer := &timerMock{}
	w := &bytes.Buffer{}
	r := bytes.NewBufferString("\n")

	wantDuration := 30
	StartTimer(w, r, timer)

	if timer.duration != wantDuration {
		t.Errorf("Error in StartTimer. Want: %v, Got: %v", wantDuration, timer.duration)
	}

	wantOutput := "Press any key to start the quiz\n"
	gotOutput := w.String()

	if wantOutput != gotOutput {
		t.Errorf("Error in StartTimer. Want: %v, Got: %v", wantOutput, gotOutput)
	}
}
