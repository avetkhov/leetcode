package main

func guessNumber(n int) int {
	left, right := 1, n

	for left <= right {
		mid := left + (right-left)/2

		answer := guess(mid)

		if answer == 0 {
			return mid
		}

		if answer == 1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
