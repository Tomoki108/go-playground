package experiments

import (
	"testing"
)

// faster
func loopWithDevision() []int {
	sl := make([]int, 0, 50000000)
	for i := 1; i <= 100000000; i++ {
		if i%2 == 0 {
			sl = append(sl, i)
		}
	}

	return sl
}

// slower
func loopWithFlag() []int {
	sl := make([]int, 0, 50000000)
	isEven := false
	for i := 1; i <= 100000000; i++ {
		if isEven {
			sl = append(sl, i)
		}

		isEven = !isEven
	}

	return sl
}

func BenchmarkLoopWithDevision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loopWithDevision()
	}
}

func BenchmarkLoopWithFlag(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loopWithFlag()
	}
}
