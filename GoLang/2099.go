package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxSubsequence([]int{2, 1, 3, 3}, 2))
}
func maxSubsequence(nums []int, k int) []int {
	dupe := nums
	sort.Ints(dupe)

	seen := make(map[int]int, k)

	for _, val := range dupe[len(dupe)-k:] {
		seen[val]++
	}

	ans := make([]int, 0, k)

	for _, val := range nums {
		if v, ok := seen[val]; ok && v > 0 {
			ans = append(ans, val)
			seen[val]--
		}
	}

	return ans
}
