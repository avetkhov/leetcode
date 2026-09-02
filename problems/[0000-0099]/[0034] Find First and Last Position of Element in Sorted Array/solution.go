package main

func searchRange(nums []int, target int) []int {
	return []int{searchFirst(nums, target), searchLast(nums, target)}
}

func searchFirst(nums []int, target int) int {
	l, r := 0, len(nums)-1
	res := -1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			res = m
			r = m - 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return res
}

func searchLast(nums []int, target int) int {
	l, r := 0, len(nums)-1
	res := -1

	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			res = m
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return res
}
