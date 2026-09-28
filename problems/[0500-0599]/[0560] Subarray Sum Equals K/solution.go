package main

func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0

	prefix := make(map[int]int)
	prefix[0] = 1

	for _, num := range nums {
		sum += num

		if freq, exists := prefix[sum-k]; exists {
			count += freq
		}

		prefix[sum]++
	}

	return count
}
