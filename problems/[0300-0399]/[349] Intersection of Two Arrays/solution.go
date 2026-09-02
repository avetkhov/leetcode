package main

func intersection(nums1 []int, nums2 []int) []int {
	s := make(map[int]struct{}, len(nums1))
	for _, x := range nums1 {
		s[x] = struct{}{}
	}

	var r []int
	for _, x := range nums2 {
		if _, f := s[x]; f {
			r = append(r, x)
			delete(s, x)
		}
	}

	return a
}
