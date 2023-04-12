package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputNum int
var groupStrNum int
var inputStr string
var check map[rune]int = make(map[rune]int)
var flag bool
var reader *bufio.Reader

func main() {
	fmt.Scanf("%d", &inputNum)
	reader = bufio.NewReader(os.Stdin)

	for i := 1; i <= inputNum; i++ {
		inputStr, _ = reader.ReadString('\n')
		flag = true
		for idx, r := range inputStr {
			if _, ok := check[r]; !ok {
				check[r] = idx
				continue
			} else if idx-check[r] > 1 {
				flag = false
				break
			}
			check[r] = idx
		}

		if flag {
			groupStrNum++
		}
		check = make(map[rune]int)
	}
	fmt.Println(groupStrNum)
}
