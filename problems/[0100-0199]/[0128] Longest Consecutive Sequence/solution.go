package main

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}

	maxLength := 0

	for num := range set {
		if _, ok := set[num-1]; ok {
			continue
		}

		length := 1

		cur := num
		for {
			if _, exists := set[cur+1]; !exists {
				break
			}
			cur++
			length++
		}

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
