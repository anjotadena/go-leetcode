package groupanagrams

import "fmt"

func groupAnagrams(strs []string) [][]string {
	result := [][]string{}

	// a, b, c, ..., z
	hashMapChars := map[[26]int][]string{}

	for i := 0; i < len(strs); i++ {
		chars := [26]int{}

		for _, c := range strs[i] {
			chars[c-'a']++
		}

		hashMapChars[chars] = append(hashMapChars[chars], strs[i])
	}

	// extract values
	for _, v := range hashMapChars {
		result = append(result, v)
	}

	return result
}

func main() {
	//
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagrams(strs))
}
