package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextGreater := make(map[int]int, len(nums2))

	stack := make([]int, 0, len(nums2))
	for i := len(nums2) - 1; i >= 0; i-- {
		curr := nums2[i]

		for len(stack) > 0 && curr > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			nextGreater[curr] = stack[len(stack)-1]
		} else {
			nextGreater[curr] = -1
		}

		stack = append(stack, curr)
	}

	var result []int
	for _, num := range nums1 {
		result = append(result, nextGreater[num])
	}

	return result
}
