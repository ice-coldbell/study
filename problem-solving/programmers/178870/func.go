package main

func solution(sequence []int, k int) []int {
	var startIDX, endIDX int
	var result [2]int
	sequencelen := len(sequence)
	length := sequencelen + 1
	sum := sequence[0]
	for {
		if sum == k && length > (endIDX-startIDX+1) {
			length = endIDX - startIDX + 1
			result[0] = startIDX
			result[1] = endIDX
		}

		switch {
		case sum > k || sum == k:
			sum -= sequence[startIDX]
			startIDX++
		case sum < k:
			endIDX++
			if endIDX == sequencelen {
				return result[:]
			}
			sum += sequence[endIDX]
		}
		if startIDX == sequencelen {
			break
		}
	}

	return result[:]
}
