package main

func pivotIndex(nums []int) int {
	total := 0
	for _, x := range nums {
		total += x
	}

	prefix := 0
	for i, x := range nums {
		if prefix == total-prefix-x {
			return i
		}

		prefix += x
	}

	return -1
}
