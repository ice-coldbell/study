package main

import (
	"sort"
	"strconv"
	"strings"
)

type Plan struct {
	Name      string
	StartTime int // min
	Remain    int // min
}

func solution(plans [][]string) []string {
	var (
		planArr []*Plan
		finish  []string
		remain  []*Plan
		current *Plan
		next    *Plan
	)

	for _, plan := range plans {
		startTime := strings.Split(plan[1], ":")
		hour, _ := strconv.Atoi(startTime[0])
		min, _ := strconv.Atoi(startTime[1])
		st := hour*60 + min
		du, _ := strconv.Atoi(plan[2])
		planArr = append(planArr, &Plan{plan[0], st, du})
	}

	sort.Slice(planArr, func(a, b int) bool {
		return planArr[a].StartTime < planArr[b].StartTime
	})

	current = planArr[0]
	next = planArr[1]
	nextIdx := 1

	for {
		current.Remain -= next.StartTime - current.StartTime
		var extraTime int
		if current.Remain <= 0 {
			extraTime = -current.Remain
			current.Remain = 0
			finish = append(finish, current.Name)
		} else {
			remain = append(remain, current)
		}

		if extraTime > 0 && len(remain) > 0 {
			for i := len(remain) - 1; i >= 0; i-- {
				breakFlag := false
				switch {
				case extraTime > remain[i].Remain:
					extraTime -= remain[i].Remain
					finish = append(finish, remain[i].Name)
					remain = append(remain[:i], remain[i+1:]...)
				case extraTime == remain[i].Remain:
					finish = append(finish, remain[i].Name)
					remain = append(remain[:i], remain[i+1:]...)
					breakFlag = true
				case extraTime < remain[i].Remain:
					remain[i].Remain -= extraTime
					breakFlag = true
				}
				if breakFlag {
					break
				}
			}
		}

		current = next
		nextIdx++
		if nextIdx == len(planArr) {
			finish = append(finish, current.Name)
			for i := len(remain) - 1; i >= 0; i-- {
				finish = append(finish, remain[i].Name)
			}
			break
		}
		next = planArr[nextIdx]
	}
	return finish
}
