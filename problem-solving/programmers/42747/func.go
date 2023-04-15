package main

import (
	"sort"
)

func solution(citations []int) int {
	sort.Ints(citations)
	length := len(citations)
	var hIndex int
	for i := 0; i < length; i++ {
		if citations[i] <= len(citations[i:]) && hIndex < citations[i] {
			hIndex = citations[i]
		} else if citations[i] > len(citations[i:]) && hIndex < len(citations[i:]) {
			hIndex = len(citations[i:])
		}
	}
	return hIndex
}
