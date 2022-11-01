package twosumiiinputarrayissorted

import "fmt"

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		lv, rv := numbers[l], numbers[r]

		// sum
		lrs := lv + rv

		if lrs == target {
			return []int{l + 1, r + 1}
		}

		if lrs < target {
			l++
		} else {
			r--
		}
	}

	return []int{-1}
}

func main() {
	input, target := []int{2, 7, 11, 15}, 9

	fmt.Println(twoSum(input, target))
}
