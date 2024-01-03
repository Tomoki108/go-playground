package main

import (
	"fmt"
	"math"
	"strconv"
)

func getSqrt(i int) int {
	return int(math.Sqrt(float64(i)))
}

func main() {
	fmt.Print("入力値:")
	var input string
	fmt.Scanln(&input)

	inputN, err := strconv.Atoi(input)
	if err != nil {
		fmt.Print("入力値が不正です")
		return
	}

	if isPrimeNumV3(inputN) {
		fmt.Printf("%dは素数です\n", inputN)
	} else {
		fmt.Printf("%dは素数ではありません\n", inputN)
	}
}
