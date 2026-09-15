package main

func maxArea(height []int) int {
	maxArea := 0

	left, right := 0, len(height)-1
	for left < right {
		curArea := min(height[left], height[right]) * (right - left)
		if curArea > maxArea {
			maxArea = curArea
		}

		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}

	return maxArea
}
