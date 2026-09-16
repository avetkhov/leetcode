package main

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x/2

	for left <= right {
		mid := left + (right-left)/2

		if mid == x/mid {
			return mid
		}

		if mid > x/mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right
}
