package main

func solution(numbers []int, target int) int {
	return dfs(numbers, target, 0, 0)
}

func dfs(numbers []int, target int, curr, idx int) int {
	if idx == len(numbers) {
		if curr == target {
			return 1
		}
		return 0
	}

	return dfs(numbers, target, curr+numbers[idx], idx+1) +
		dfs(numbers, target, curr-numbers[idx], idx+1)
}
