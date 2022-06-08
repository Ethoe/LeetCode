package problems

import "fmt"

func (p Problem) Run88() {
	fmt.Println("Problem 88: https://leetcode.com/problems/merge-sorted-array/")
	merge([]int{1,2,3}, 3, []int{2,5,6}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int)  {   
    i, j, k := m-1, n-1, m+n-1
    for ; j>=0; k-- {
        if i >=0 && nums1[i] > nums2[j] {
            nums1[k] = nums1[i]
            i--
        } else {
            nums1[k] = nums2[j]
            j--
        }
    }
}