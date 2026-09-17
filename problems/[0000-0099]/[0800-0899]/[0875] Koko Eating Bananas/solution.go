package main

func minEatingSpeed(piles []int, h int) int {
	maxPile := 0

	for _, pile := range piles {
		if pile > maxPile {
			maxPile = pile
		}
	}

	left, right := 1, maxPile

	for left <= right {
		mid := left + (right-left)/2

		x := 0
		for _, pile := range piles {
			x += (pile + mid - 1) / mid
		}

		if x <= h {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}
