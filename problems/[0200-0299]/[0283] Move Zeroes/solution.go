package main

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}

	left := 0

	for right, num := range nums {
		if num != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
