package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	num1Len, i := len(nums1), 0
	num2Len, j := len(nums2), 0

	// initialiaze slice with a size of num1 length plus num2 length
	result := make([]int, num1Len+num2Len)

	for z := 0; z < (num1Len + num2Len); z++ {
		if i == num1Len || (i < num1Len && j < num2Len && nums1[i] > nums2[j]) {
			result[z] = nums2[j]
			j++
			continue
		}

		if j == num2Len || (i < num1Len && j < num2Len && nums1[i] <= nums2[j]) {
			result[z] = nums1[i]
			i++
		}
	}

	if len(result) == 0 {
		panic("Error!")
	}

	if len(result)%2 == 0 {
		return float64((result[len(result)/2] + result[len(result)/2-1])) / 2.0
	}

	return float64(result[len(result)/2])
}

func main() {
	//
}
