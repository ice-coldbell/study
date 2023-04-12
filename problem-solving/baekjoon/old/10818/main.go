package main

import (
	"bufio"
	"fmt"
	"os"
)

var a, min, max int
var c byte
var op bool
var reader *bufio.Reader

func main() {
	reader = bufio.NewReaderSize(os.Stdin, 15000000)
	reader.ReadBytes('\n')
	input, _ := reader.ReadBytes('\n')

	min, max = 1000000, -1000000
	op = false
	for _, c = range input {
		switch {
		case c >= '0' && c <= '9':
			a = a*10 + int(c-'0')
		case c == '-':
			op = true
		default:
			if op {
				a *= -1
			}
			if max < a {
				max = a
			}
			if min > a {
				min = a
			}
			op = false
			a = 0
		}
	}
	fmt.Printf("%d %d", min, max)
}
