package main

import "fmt"

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		// add l to avoid overflowing
		m := (l + (r-l)/2)

		if nums[m] == target {
			return m
		}

		if target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	r := search(nums, target)

	fmt.Printf("nums: %v | target: %d | result: %d\n", nums, target, r)
}
