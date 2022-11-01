package movezeroes

func moveZeroes(nums []int) {
	l := 0

	for r := range nums {
		if nums[r] != 0 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}
}

func main() {
	n := []int{0, 1, 0, 3, 12}

	moveZeroes(n)
}
