package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testcases = []struct {
	plans  [][]string
	result []string
}{
	{
		[][]string{{"korean", "11:40", "30"}, {"english", "12:10", "20"}, {"math", "12:30", "40"}},
		[]string{"korean", "english", "math"},
	},
	{
		[][]string{{"science", "12:40", "50"}, {"music", "12:20", "40"}, {"history", "14:00", "30"}, {"computer", "12:30", "100"}},
		[]string{"science", "history", "computer", "music"},
	},
	{
		[][]string{{"aaa", "12:00", "20"}, {"bbb", "12:10", "30"}, {"ccc", "12:40", "10"}},
		[]string{"bbb", "ccc", "aaa"},
	},
}

func TestXxx(t *testing.T) {
	for i, tc := range testcases {
		tc := tc
		t.Run(strconv.Itoa(i),
			func(t *testing.T) {
				result := solution(tc.plans)
				assert.Equal(t, tc.result, result)
			})
	}
}
