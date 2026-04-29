package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	res := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]

			if sum == target {
				return sum
			}

			if abs(target-sum) < abs(target-res) {
				res = sum
			}

			if sum > target {
				r--
			} else {
				l++
			}
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
