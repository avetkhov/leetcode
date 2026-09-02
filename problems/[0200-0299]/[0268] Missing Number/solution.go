package main

func missingNumber(nums []int) int {
	s := len(nums) * (len(nums) + 1) / 2

	for _, x := range nums {
		s -= x
	}

	return s
}
