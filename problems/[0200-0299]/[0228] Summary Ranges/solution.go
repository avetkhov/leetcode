package main

import "strconv"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	var result []string

	for left, right := 0, 0; right < len(nums); right++ {
		if right == len(nums)-1 || nums[right+1] != nums[right]+1 {
			if left == right {
				result = append(result, strconv.Itoa(nums[left]))
			} else {
				result = append(result, strconv.Itoa(nums[left])+"->"+strconv.Itoa(nums[right]))
			}
			left = right + 1
		}
	}

	return result
}
