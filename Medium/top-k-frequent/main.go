package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	r := make([]int, 0, k)

	cMap := make(map[int]int, len(nums))

	for _, n := range nums {
		cMap[n]++
	}

	counts := make([]int, 0, len(cMap))

	for _, c := range cMap {
		counts = append(counts, c)
	}

	sort.Ints(counts)

	min := counts[len(counts)-k]

	for n, c := range cMap {
		if c >= min {
			r = append(r, n)
		}
	}

	return r
}

func main() {
	k := 2
	nums := []int{1, 1, 1, 2, 2, 3}

	topK := topKFrequent(nums, k)

	fmt.Printf("Top K Frequent: %v", topK)
}
