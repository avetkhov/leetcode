package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	left := 0

	for _, num := range nums {
		if num != nums[left] {
			left++
			nums[left] = num
		}
	}

	return left + 1
}
