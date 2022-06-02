package problems

import "fmt"

func Run187() {
	fmt.Println("Problem 187: https://leetcode.com/problems/repeated-dna-sequences/")
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	fmt.Println(findRepeatedDnaSequences(s))
}

func findRepeatedDnaSequences(s string) []string {
	window := 10
	dnaMap := make(map[string]int)
	for sequence := 0; sequence+window <= len(s); sequence++ {
		substring := s[sequence : sequence+window]
		dnaMap[substring]++
	}
	res := make([]string, 0)
	for key, value := range dnaMap {
		if value > 1 {
			res = append(res, key)
		}
	}
	return res
}
