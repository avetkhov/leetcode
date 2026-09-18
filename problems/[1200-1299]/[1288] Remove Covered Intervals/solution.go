package main

import "sort"

func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		return intervals[i][1] > intervals[j][1]
	})

	count := len(intervals)
	end := 0

	for _, interval := range intervals {
		if interval[1] > end {
			end = interval[1]
		} else {
			count--
		}
	}

	return count
}
