package main

func intersect(nums1 []int, nums2 []int) []int {
	s := make(map[int]int, len(nums1))
	for _, x := range nums1 {
		s[x]++
	}

	r := make([]int, 0, len(nums1))
	for _, x := range nums2 {
		if s[x] > 0 {
			r = append(r, x)
			s[x]--
		}
	}

	return r
}
