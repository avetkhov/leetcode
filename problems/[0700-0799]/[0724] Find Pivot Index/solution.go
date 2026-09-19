package main

func pivotIndex(nums []int) int {
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

func pivotIndex(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	prefix := make([]int, len(nums)+1)
	for i, num := range nums {
		if totalSum-num == 2*prefix[i] {
			return i
		}

		prefix[i+1] = prefix[i] + num
	}

	return -1
}
