package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testcases = []struct {
	array    []int
	commands [][]int
	result   []int
}{
	{
		array: []int{1, 5, 2, 6, 3, 7, 4},
		commands: [][]int{
			{2, 5, 3},
			{4, 4, 1},
			{1, 7, 3},
		},
		result: []int{5, 6, 3},
	},
}

func TestXxx(t *testing.T) {
	for i, tc := range testcases {
		t.Run(strconv.Itoa(i),
			func(t *testing.T) {
				result := solution(tc.array, tc.commands)
				assert.Equal(t, tc.result, result)
			})
	}
}
