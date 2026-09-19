package main

func nextGreaterElements(nums []int) []int {
	result := make([]int, len(nums))

	for i := range result {
		result[i] = -1
	}

	stack := make([]int, 0, len(nums))
	for i := 2*len(nums) - 1; i >= 0; i-- {
		idx := i % len(nums)

		for len(stack) > 0 && nums[idx] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			result[idx] = nums[stack[len(stack)-1]]
		}

		stack = append(stack, idx)
	}

	return result
}

func nextGreaterElements(nums []int) []int {
	result := make([]int, len(nums))

	for i := range result {
		result[i] = -1
	}

	stack := make([]int, 0, len(nums))
	for i := 0; i < 2*len(nums); i++ {
		idx := i % len(nums)

		for len(stack) > 0 && nums[idx] > nums[stack[len(stack)-1]] {
			result[stack[len(stack)-1]] = nums[idx]
			stack = stack[:len(stack)-1]
		}

		if i < len(nums) {
			stack = append(stack, idx)
		}
	}

	return result
}

func main() {
	nums := []int{5, 4, 3, 2, 1}
	nextGreaterElements(nums)
}
