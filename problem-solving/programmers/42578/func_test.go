package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testcases = []struct {
	clothes [][]string
	result  int
}{
	{[][]string{{"yellow_hat", "headgear"}, {"blue_sunglasses", "eyewear"}, {"green_turban", "headgear"}}, 5},
	{[][]string{{"crow_mask", "face"}, {"blue_sunglasses", "face"}, {"smoky_makeup", "face"}}, 3},
}

func TestXxx(t *testing.T) {
	for i, tc := range testcases {
		t.Run(strconv.Itoa(i),
			func(t *testing.T) {
				result := solution(tc.clothes)
				assert.Equal(t, tc.result, result)
			})
	}
}
