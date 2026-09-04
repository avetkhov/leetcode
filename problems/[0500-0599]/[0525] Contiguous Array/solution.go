package main

func findMaxLength(nums []int) int {
	maxLength := 0
	sum := 0
	pm := make(map[int]int, len(nums)+1)
	pm[0] = -1

	for i, num := range nums {
		if num == 1 {
			sum++
		} else {
			sum--
		}

		if j, ok := pm[sum]; ok {
			length := i - j
			if length > maxLength {
				maxLength = length
			}
		} else {
			pm[sum] = i
		}
	}

	return maxLength
}
