package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testcases = []struct {
	numbers []int
	target  int
	result  int
}{
	{[]int{1, 1, 1, 1, 1}, 3, 5},
	{[]int{4, 1, 2, 1}, 4, 2},
}

func TestXxx(t *testing.T) {
	for i, tc := range testcases {
		t.Run(strconv.Itoa(i),
			func(t *testing.T) {
				result := solution(tc.numbers, tc.target)
				assert.Equal(t, tc.result, result)
			})
	}
}
