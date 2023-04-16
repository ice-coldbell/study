package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testcases = []struct {
	numbers string
	result  int
}{
	{"17", 3},
	{"011", 2},
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
