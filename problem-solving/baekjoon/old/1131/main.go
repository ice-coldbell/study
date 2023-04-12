package main

import (
	"fmt"
	"math"
)

var pow_array map[int]int
var d []int
var k int
var cash map[int]map[int]struct{} = map[int]map[int]struct{}{
	1: {1: {}, 2: {}, 3: {}, 4: {}, 5: {}, 6: {}, 7: {}, 8: {}, 9: {}},
	2: {1: {}, 4: {}},
	3: {1: {}, 55: {}, 136: {}, 153: {}, 160: {}, 370: {}, 371: {}, 407: {}, 919: {}},
	4: {1: {}, 1138: {}, 1634: {}, 2178: {}, 8208: {}, 9474: {}},
	5: {1: {}, 244: {}, 4150: {}, 4151: {}, 8294: {}, 8299: {}, 9044: {}, 9045: {}, 10933: {}, 24584: {}, 54748: {}, 58618: {}, 89883: {}, 92727: {}, 93084: {}, 194979: {}},
	6: {1: {}, 17148: {}, 63804: {}, 93531: {}, 239459: {}, 282595: {}, 548834: {}},
}

func init() {
	pow_array = make(map[int]int)
	d = make([]int, 1000001)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b, &k)
	for i := 0; i < 10; i++ {
		pow_array[i] = int(math.Pow(float64(i), float64(k)))
	}

	for key := range cash[k] {
		n := key
		for {
			if n < 1000000 {
				d[n] = key
			}

			n = S(n)

			if n == key {
				break
			}
		}
	}

	var sum int64
	for i := a; i <= b; i++ {
		sum += int64(min(i))
	}
	fmt.Println(sum)
}

func min(n int) int {
	for {
		if n > 1000000 {
			return min(S(n))
		}

		if d[n] > 0 {
			return d[n]
		}

		t := min(S(n))
		if t > n {
			d[n] = n
		} else {
			d[n] = t
		}
		return min(d[n])
	}
}

func S(n int) (sum int) {
	for n != 0 {
		sum += pow_array[n%10]
		n /= 10
	}
	return sum
}
