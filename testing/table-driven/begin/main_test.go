// testing/table-driven/begin/main_test.go
package main

import (
	
	"testing"
)

func TestSum(t *testing.T) {
	var tests = []struct {
		name  string
		input []int
		want  int
	}{
		{"one", []int{1}, 1},
		{"two", []int{1, -2}, -1},
		{"three", []int{1, 2, 3}, 6},
	}

	// range over the tests and run them as subtests
	//
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T){
			got := sum (testCase.input...)
			if got != testCase.want{
				t.Errorf("Test Failed %v %v", got, testCase.want )
			}
		})
	}
}
