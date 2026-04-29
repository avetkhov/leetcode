package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i, x := range nums {
		if x > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for l, r := i+1, len(nums)-1; l < r; {
			if l > i+1 && nums[l] == nums[l-1] {
				l++
				continue
			}

			sum := nums[i] + nums[l] + nums[r]

			if sum < 0 {
				l++
				continue
			}

			if sum > 0 {
				r--
				continue
			}

			res = append(res, []int{nums[i], nums[l], nums[r]})
			l++
			r--
		}
	}

	return res
}
