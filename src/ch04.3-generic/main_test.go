package main

import "testing"

func BenchmarkBuiltinTypes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumInts(ints)
		SumFloats(floats)
	}
}

func BenchmarkGeneric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumNumbers(ints)
		SumNumbers(floats)
	}
}
