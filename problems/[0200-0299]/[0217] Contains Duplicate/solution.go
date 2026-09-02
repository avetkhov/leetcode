package main

func containsDuplicate(nums []int) bool {
	d := map[int]int{}
	var r []int
	for i, num := range nums {
		if _, found := d[num]; found {
			return found
		}
		d[num] = i
	}

	return false
}
