package main

import "sort"

func task56(intervals [][]int) [][]int {
	arrs := make([][]int, len(intervals))
	if len(intervals) == 0 {
		return arrs
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	i := 0
	res := [][]int{}
	for i < len(intervals) {
		l := intervals[i][0]
		r := intervals[i][1]

		var j int
		for j = i + 1; j < len(intervals); j++ {
			if intervals[j][0] <= r {
				r = max(r, intervals[j][1])
			} else {
				break
			}
		}
		res = append(res, []int{l, r})
		i = j
	}
	return res

}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
