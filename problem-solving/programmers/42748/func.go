package main

import (
	"sort"
)

func solution(array []int, commands [][]int) []int {
	result := make([]int, 0, len(commands))
	for _, cmd := range commands {
		result = append(result, kNum(array, cmd))
	}
	return result
}

func kNum(src []int, command []int) int {
	arr := make([]int, 0, len(src))
	for i := range src {
		arr = append(arr, src[i])
	}

	var (
		i = command[0]
		j = command[1]
		k = command[2]
	)

	arr = arr[i-1 : j]
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	return arr[k-1]
}
