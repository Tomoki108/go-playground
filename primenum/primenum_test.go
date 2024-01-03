package main

import (
	"testing"
)

// 3rd fastest
func BenchmarkIsPrimeNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrimeNum(99563)
	}
}

// 2nd fastest
func BenchmarkIsPrimeNumV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrimeNumV2(99563)
	}
}

// fastest but memory consuming because of recursive function call
func BenchmarkIsPrimeNumV3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrimeNumV3(99563)
	}
}
