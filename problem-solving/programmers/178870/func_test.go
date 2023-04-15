package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testcases = []struct {
	sequence []int
	k        int
	result   []int
}{
	{
		sequence: []int{1, 2, 3, 4, 5},
		k:        7,
		result:   []int{2, 3},
	},
	{
		sequence: []int{1, 1, 1, 2, 3, 4, 5},
		k:        5,
		result:   []int{6, 6},
	},
	{
		sequence: []int{2, 2, 2, 2, 2},
		k:        6,
		result:   []int{0, 2},
	},
	{
		sequence: []int{6, 2, 2, 2, 2},
		k:        6,
		result:   []int{0, 0},
	},
	{
		sequence: []int{1, 2, 6, 2, 2},
		k:        6,
		result:   []int{2, 2},
	},
	{
		sequence: []int{1, 6, 6, 2, 2},
		k:        6,
		result:   []int{1, 1},
	},
	{
		sequence: []int{2, 2, 2, 2, 2},
		k:        10,
		result:   []int{0, 4},
	},
}

func TestXxx(t *testing.T) {
	for i, tc := range testcases {
		t.Run(strconv.Itoa(i),
			func(t *testing.T) {
				result := solution(tc.sequence, tc.k)
				assert.Equal(t, tc.result, result)
			})
	}
}
