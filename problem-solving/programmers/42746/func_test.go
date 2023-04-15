package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testcases = []struct {
	numbers []int
	result  string
}{
	{[]int{6, 10, 2}, "6210"},
	{[]int{3, 30, 34, 5, 9}, "9534330"},
	{[]int{1, 20, 300, 4000, 50000}, "500004000300201"},
	{[]int{0, 0, 0, 0, 0, 0, 10, 0}, "100000000"},
	{[]int{0, 0, 0, 0, 0, 0, 0}, "0"},
}

func TestXxx(t *testing.T) {
	for i, tc := range testcases {
		t.Run(strconv.Itoa(i),
			func(t *testing.T) {
				result := solution(tc.numbers)
				assert.Equal(t, tc.result, result)
			})
	}
}
