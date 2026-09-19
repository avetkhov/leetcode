package main

func sumOddLengthSubarrays(arr []int) int {
	prefix := make([]int, len(arr)+1)
	for i, num := range arr {
		prefix[i+1] = prefix[i] + num
	}

	sum := 0
	for i := 1; i <= len(arr); i = i + 2 {
		for j := 0; i+j < len(prefix); j++ {
			sum += prefix[i+j] - prefix[j]
		}
	}

	return sum
}
