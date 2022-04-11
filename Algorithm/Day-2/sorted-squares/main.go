package main

func sortedSquares(nums []int) []int {
	l := 0
	p := len(nums) - 1
	r := p

	ans := make([]int, p+1)

	for l <= r {
		lv := nums[l] * nums[l]
		rv := nums[r] * nums[r]

		if lv > rv {
			ans[p] = lv
			p = p - 1
			l = l + 1
		} else {
			ans[p] = rv
			p = p - 1
			r = r - 1
		}
	}

	return ans
}

func main() {}
