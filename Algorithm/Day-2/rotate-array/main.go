package main

func rotate(nums []int, k int) {
	// pointer
	k %= len(nums)

	// Reverse original array
	l, r := 0, len(nums)-1
	reverseValues(&nums, l, r)

	// Reverse original array from 0 index to k
	l, r = 0, k-1
	reverseValues(&nums, l, r)

	// Reverse original array from k to size of array
	l, r = k, len(nums)-1
	reverseValues(&nums, l, r)
}

func reverseValues(n *[]int, l int, r int) {
	for l < r {
		(*n)[l], (*n)[r] = (*n)[r], (*n)[l]
		l++
		r--
	}
}

func main() {
	// input
	nums, k := []int{1, 2, 3, 4, 5, 6, 7}, 3

	// solve
	rotate(nums, k)
}
