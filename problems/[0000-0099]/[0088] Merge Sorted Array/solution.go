package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	for left1, right2, right1 := m-1, n-1, m+n-1; right2 >= 0; right1-- {
		if left1 >= 0 && nums1[left1] > nums2[right2] {
			nums1[right1] = nums1[left1]
			left1--
		} else {
			nums1[right1] = nums2[right2]
			right2--
		}
	}
}
