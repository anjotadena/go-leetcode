package main

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := high + (low-high)/2

		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}

func main() {

}
