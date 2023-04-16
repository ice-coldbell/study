package main

import (
	"fmt"
	"math"
	"strconv"
)

func solution(numbers string) int {
	var result int
	dfs(numbers, "entry", &result)
	return result
}

func dfs(numbers string, number string, result *int) {
	if number != "entry" {
		INumber, _ := strconv.Atoi(number)
		fmt.Println(INumber)
		if checkPrimeNumber(INumber) {
			*result += 1
		}
	}

	if len(numbers) == 0 {
		return
	}

	for i, r := range numbers {
		src := string(append([]rune(numbers[:i]), []rune(numbers[i+1:])...))
		if number == "entry" {
			number = ""
		}
		nextNum := number + string(r)
		dfs(src, nextNum, result)
	}
}

func checkPrimeNumber(num int) bool {
	if num == 1 || num == 0 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
