package main

func minSubArrayLen(target int, nums []int) int {
	minLength := len(nums) + 1
	left, sum := 0, 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			length := right - left + 1

			if length < minLength {
				minLength = length
			}

			sum -= nums[left]
			left++
		}
	}

	if minLength == len(nums)+1 {
		return 0
	}

	return minLength
}
