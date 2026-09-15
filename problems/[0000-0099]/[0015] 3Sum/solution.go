package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	// quickSort(nums)

	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1
		for left < right {
			if left > i+1 && nums[left] == nums[left-1] {
				left++
				continue
			}

			sum := nums[i] + nums[left] + nums[right]

			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			}
		}
	}

	return result
}

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2
	nums[mid], nums[len(nums)-1] = nums[len(nums)-1], nums[mid]

	left := 0
	for right := 0; right < len(nums)-1; right++ {
		if nums[right] < nums[len(nums)-1] {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}

	nums[left], nums[len(nums)-1] = nums[len(nums)-1], nums[left]

	quickSort(nums[:left])
	quickSort(nums[left+1:])

	return nums
}
