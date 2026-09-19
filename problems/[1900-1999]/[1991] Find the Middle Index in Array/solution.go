package main

func findMiddleIndex(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	prefixSum := 0
	for i, num := range nums {
		if totalSum-num == 2*prefixSum {
			return i
		}

		prefixSum += num
	}

	return -1
}
