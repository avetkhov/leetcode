package main

func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int, len(nums))

	for i, num := range nums {
		if j, ok := seen[num]; ok && i-j <= k {
			return true
		}
		seen[num] = i
	}

	return false
}
