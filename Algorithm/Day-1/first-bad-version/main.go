package main

import "fmt"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func isBadVersion(n int) bool {
	return true
}

func firstBadVersion(n int) int {
	low := 1
	high := n
	result := n
	mid := n / 2

	for low <= high {
		mid = high + ((low - high) / 2)

		if isBadVersion(mid) {
			result = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return result
}

func main() {
	fmt.Printf("Bad version: %b\n", firstBadVersion(4))
}
