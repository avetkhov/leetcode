package main

func longestOnes(nums []int, k int) int {
	l, zeroCount, maxLength := 0, 0, 0

	for r := 0; r < len(nums); r++ {
		if nums[r] == 0 {
			zeroCount++
		}

		for zeroCount > k {
			if nums[l] == 0 {
				zeroCount--
			}
			l++
		}

		length := r - l + 1
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
