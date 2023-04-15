package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testcases = []struct {
	result  int
	n       int
	lost    []int
	reserve []int
}{
	{5, 5, []int{2, 4}, []int{1, 3, 5}},
	{4, 5, []int{2, 4}, []int{3}},
	{2, 3, []int{3}, []int{1}},
	{3, 5, []int{1, 5}, []int{3}},
	{3, 5, []int{1, 3, 4, 5}, []int{3, 4}},
	{0, 5, []int{1, 2, 3, 4, 5}, []int{}},
	{5, 5, []int{}, []int{1, 2, 3, 4, 5}},
	{4, 4, []int{1, 3}, []int{2, 4}},
	{4, 4, []int{2, 4}, []int{1, 3}},
	{6, 7, []int{1, 3, 4, 6}, []int{2, 5, 7}},
	{6, 7, []int{1, 3, 4, 6}, []int{2, 5, 7}},
}

func TestXxx(t *testing.T) {
	for i, tc := range testcases {
		t.Run(strconv.Itoa(i),
			func(t *testing.T) {
				result := solution(tc.n, tc.lost, tc.reserve)
				assert.Equal(t, tc.result, result)
			})
	}
}
