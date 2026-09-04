package main

func containsNearbyDuplicate(nums []int, k int) bool {
	s := make(map[int]int, len(nums))

	for i, v := range nums {
		if j, ok := s[v]; ok && i-j <= k {
			return true
		}
		s[v] = i
	}

	return false
}
