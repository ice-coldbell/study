package main

import (
	"fmt"
)

var (
	n                int
	startNum, endNum int = 1, 1
	cnt              int = 1
	sum              int = 1
)

func main() {
	fmt.Scan(&n)
	for startNum != n && endNum != n {
		switch {
		case sum == n:
			cnt++
			endNum++
			sum += endNum - startNum
			startNum++
		case sum < n:
			endNum++
			sum += endNum
		case sum > n:
			sum -= startNum
			startNum++
		}
	}
	fmt.Printf("%d", cnt)
}
