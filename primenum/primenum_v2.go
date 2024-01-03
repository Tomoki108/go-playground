package main

import (
	"fmt"
)

func isPrimeNumV2(n int) bool {
	if n < 0 || n == 0 || n == 1 {
		return false
	}
	if n == 2 {
		return true
	}

	sqrtn := getSqrt(n)

	factorCandidates := make([]int, sqrtn/2+1)
	checkList := make(map[int]bool, sqrtn/2+1)
	factorCandidates[0] = 2
	checkList[2] = false
	for i := 3; i <= sqrtn; i++ {
		if i%2 == 0 {
			continue
		}
		factorCandidates[i/2] = i
		checkList[i] = false
	}

	for _, fc := range factorCandidates {
		if checkList[fc] {
			continue
		}

		if n%fc == 0 {
			fmt.Printf("%dは%dで割り切れます\n", n, fc)
			return false
		}

		// 因数候補の倍数を検証済みにする
		fcs := fc * fc
		for fc2, checked2 := range checkList {
			if fc2 < fcs || checked2 {
				continue
			}
			if fc2%fc == 0 {
				checkList[fc2] = true
			}
		}
	}

	return true
}
