package main

import (
	"sort"
	"strconv"
)

func solution(numbers []int) string {
	sort.Slice(numbers,
		func(i, j int) bool {
			iStr := strconv.Itoa(numbers[i])
			jStr := strconv.Itoa(numbers[j])
			iFirst := iStr + jStr
			jFirst := jStr + iStr
			a, _ := strconv.Atoi(iFirst)
			b, _ := strconv.Atoi(jFirst)

			return a > b
		},
	)
	var result string
	for _, num := range numbers {
		result += strconv.Itoa(num)
	}

	a, _ := strconv.Atoi(result)
	if a == 0 {
		return "0"
	}

	return result
}
