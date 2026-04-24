package main

func maxArea(height []int) int {
	s := 0
	l, r := 0, len(height)-1

	for l < r {
		t := (r - l) * min(height[l], height[r])

		if s < t {
			s = t
		}

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return s
}
