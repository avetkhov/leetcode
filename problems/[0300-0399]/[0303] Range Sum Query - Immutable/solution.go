package main

type NumArray struct {
	prefix []int
}

func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)
	for i, x := range nums {
		prefix[i+1] = prefix[i] + x
	}

	return NumArray{prefix: prefix}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.prefix[right+1] - this.prefix[left]
}
