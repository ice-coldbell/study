package main

import (
	"sort"
)

func solution(n int, lost []int, reserve []int) int {
	m := make(map[int]struct{}, len(lost))
	for _, l := range lost {
		m[l] = struct{}{}
	}

	pass := map[int]struct{}{}
	for _, num := range reserve {
		if _, ok := m[num]; ok {
			delete(m, num)
			pass[num] = struct{}{}
		}
	}

	sort.SliceStable(reserve, func(i, j int) bool {
		return reserve[i] < reserve[j]
	})
	for _, num := range reserve {
		if _, ok := pass[num]; ok {
			continue
		}

		if _, ok := m[num-1]; ok {
			delete(m, num-1)
			continue
		}

		if _, ok := m[num+1]; ok {
			delete(m, num+1)
			continue
		}
	}
	return n - len(m)
}
