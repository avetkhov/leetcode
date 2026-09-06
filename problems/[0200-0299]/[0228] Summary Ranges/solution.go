package main

import "strconv"

func summaryRanges(nums []int) []string {
	n := len(nums)
	res := make([]string, 0, n)

	for i := 0; i < n; i++ {
		start := nums[i]

		for i < n-1 && nums[i]+1 == nums[i+1] {
			i++
		}

		if start == nums[i] {
			res = append(res, strconv.Itoa(start))
		} else {
			res = append(res, strconv.Itoa(start)+"->"+strconv.Itoa(nums[i]))
		}
	}

	return res
}
