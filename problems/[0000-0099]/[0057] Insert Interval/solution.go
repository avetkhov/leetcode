package main

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0, len(intervals)+1)

	curStart, curEnd := newInterval[0], newInterval[1]
	i := 0

	for i < len(intervals) && intervals[i][1] < curStart {
		result = append(result, intervals[i])
		i++
	}

	for i < len(intervals) && intervals[i][0] <= curEnd {
		curStart = min(curStart, intervals[i][0])
		curEnd = max(curEnd, intervals[i][1])
		i++
	}

	result = append(result, []int{curStart, curEnd})

	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}
