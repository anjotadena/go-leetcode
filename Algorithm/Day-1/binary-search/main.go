package main

import "fmt"

func search(nums []int, target int) int {
	var low, high = 0, len(nums) - 1
	median := 0

	for low <= high {
		median = low + (high-low)/2

		if nums[median] == target {
			return median
		}

		if nums[median] < target {
			low = median + 1
		} else {
			high = median - 1
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
