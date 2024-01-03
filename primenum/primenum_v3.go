package main

import (
	"fmt"
)

func checkFromHeadTotail(factorCandidates []int, n int) bool {
	if len(factorCandidates) == 0 {
		return true
	}

	fc := factorCandidates[0]
	if n%fc == 0 {
		fmt.Printf("%dは%dで割り切れます\n", n, fc)
		return false
	}

	oldLength := len(factorCandidates)
	numOfFcMultiple := oldLength / fc
	newLength := oldLength - numOfFcMultiple

	newFactorCandidates := make([]int, 0, newLength)
	for _, fc2 := range factorCandidates {
		if fc2%fc == 0 {
			continue
		}

		newFactorCandidates = append(newFactorCandidates, fc2)
	}

	return checkFromHeadTotail(newFactorCandidates, n)
}

func isPrimeNumV3(n int) bool {
	if n < 0 || n == 0 || n == 1 {
		return false
	}
	if n == 2 {
		return true
	}

	sqrtn := getSqrt(n)
	factorCandidates := make([]int, sqrtn/2+1)
	factorCandidates[0] = 2
	for i := 3; i <= sqrtn; i++ {
		if i%2 == 0 {
			continue
		}
		factorCandidates[i/2] = i
	}

	return checkFromHeadTotail(factorCandidates, n)
}
