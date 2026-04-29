package main

func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	l := 2

	for r := 2; r < len(nums); r++ {
		if nums[r] != nums[l-2] {
			nums[l] = nums[r]
			l++
		}
	}

	return l
}
