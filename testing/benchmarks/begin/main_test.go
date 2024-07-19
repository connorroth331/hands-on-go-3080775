// testing/benchmarks/begin/main_test.go
package main

import "testing"

// write a benchmark for sum
func BenchmarkSum(b *testing.B) {
	for x := 0; x < b.N; x++{
		sum(1, 2, 3)
	}

}

func BenchmarkSumAny(b *testing.B) {
	for x := 0; x < b.N; x++{
		sumAny([]interface{}{1, 2, 3})
	}

}

// write a benchmark for sumAny
