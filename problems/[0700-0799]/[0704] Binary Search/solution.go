package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2

		if nums[m] == target {
			return m
		}

		if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}
