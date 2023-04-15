package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testcases = []struct {
	citations []int
	result    int
}{
	{[]int{3, 0, 6, 1, 5}, 3},
	{[]int{3, 0, 6, 1, 5, 7, 8, 9, 6, 6}, 6},
	{[]int{3, 0, 6, 1, 5, 7, 8, 9, 6, 6}, 6},
	{[]int{9, 7, 6, 2, 1}, 3},
}

func TestXxx(t *testing.T) {
	for i, tc := range testcases {
		t.Run(strconv.Itoa(i),
			func(t *testing.T) {
				result := solution(tc.citations)
				assert.Equal(t, tc.result, result)
			})
	}
}
