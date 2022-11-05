package longestsubstringwithoutrepeatingcharacters

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	var i, max = 0, 0

	charMap := make(map[byte]int)

	for j := 0; j < len(s); j++ {
		// every iteration increment it by 1
		charMap[s[j]]++

		for charMap[s[j]] == 2 && i < j {
			charMap[s[i]]--
			i++
		}

		if max < j-i+1 {
			max = j - i + 1
		}

	}

	return max
}
