package main

func solution(clothes [][]string) int {
	clothesMap := make(map[string]int)
	for _, arr := range clothes {
		clothesMap[arr[1]] += 1
	}

	result := 1
	for _, num := range clothesMap {
		result *= (num + 1)
	}
	result -= 1

	return result
}
