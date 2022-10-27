package longestpalindromicsubstring

import "fmt"

func longestPalindrome(s string) string {
	result := ""
	resultLength := 0

	for i := 0; i < len(s); i++ {
		// Odd length scenario
		l, r := i, i

		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > resultLength {
				result = s[l : r+1]
				resultLength = r - l + 1
			}

			l -= 1
			r += 1
		}

		l, r = i, i+1

		// Even length scenario
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if (r - l + 1) > resultLength {
				result = s[l : r+1]
				resultLength = r - l + 1
			}

			l -= 1
			r += 1
		}
	}
	return result
}

func main() {
	s := "abab"

	fmt.Println(longestPalindrome(s))
}
