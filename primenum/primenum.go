package main

import (
	"fmt"
)

func isPrimeNum(n int) bool {
	if n < 0 || n == 0 || n == 1 {
		return false
	}
	if n == 2 {
		return true
	}

	sqrtn := getSqrt(n)
	factorCandidates := make([]int, sqrtn-1)
	checkList := make(map[int]bool, sqrtn-1)
	for i := 2; i <= sqrtn; i++ {
		factorCandidates[i-2] = i
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
		for fc2, checked2 := range checkList {
			if fc2 < fc^2 || checked2 {
				continue
			}
			if fc2%fc == 0 {
				checkList[fc2] = true
			}
		}
	}

	return true
}
