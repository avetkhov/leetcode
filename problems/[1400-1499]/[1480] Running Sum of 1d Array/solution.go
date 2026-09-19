package main

func runningSum(nums []int) []int {
	prefix := make([]int, len(nums)+1)

	for i, num := range nums {
		prefix[i+1] = prefix[i] + num
	}

	return prefix[1:]
}

func runningSum(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}

	return nums
}
