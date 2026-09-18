package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0
	prevEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < prevEnd {
			count++
		} else {
			prevEnd = intervals[i][1]
		}
	}

	return count
}
