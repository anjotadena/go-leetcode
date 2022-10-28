package main

import "fmt"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
// this is just a dummy
// this function already defined in leetcode
// for local purpose we mock it as true
func isBadVersion(n int) bool {
	return true
}

func firstBadVersion(n int) int {
	var low, high, median = 1, n, n / 2

	for low < high {
		median = low + (high-low)/2

		if isBadVersion(median) {
			high = median
		} else {
			low = median + 1
		}
	}

	return low
}

func main() {
	fmt.Printf("Bad version: %b\n", firstBadVersion(4))
}
