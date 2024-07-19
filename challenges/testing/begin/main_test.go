// challenges/testing/begin/main_test.go
package main

import (
	"testing"
)

// write a test for letterCounter.count

func TestLetterCount(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  int
	}{
		{"one", "a2_32_^_&", 1},
		{"two", "812_%6ab", 2},
		{"three", "#00", 0},
	}
	l := letterCounter{}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T){
			got := l.count(testCase.input)
			if got != testCase.want{
				t.Errorf("Test Failed %v %v", got, testCase.want )
			}
		})
	}

}
// write a test for numberCounter.count

// write a test for symbolCounter.count
