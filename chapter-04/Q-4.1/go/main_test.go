package main

import (
	"testing"
)

func Benchmark_wayNumRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wayNumRecursive(10, 20)
	}
}

func Benchmark_wayNumDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wayNumDP(11, 21)
	}
}

func Benchmark_wayNumRecursiveMemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		initMemory() // initMemory() cost the most of time
		wayNumRecursiveMemo(10, 20)
	}
}
