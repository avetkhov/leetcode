package main

func twoSum(nums []int, target int) []int {
	d := map[int]int{}

	for i, num := range nums {
		x := target - num
		if index, found := d[x]; found {
			return []int{index, i}
		}
		d[num] = i
	}

	return []int{-1, -1}
}
