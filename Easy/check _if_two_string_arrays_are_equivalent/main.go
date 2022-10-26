package main

import "fmt"

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	a, b := "", ""

	for i := 0; i < len(word1); i++ {
		a += word1[i]
	}

	for i := 0; i < len(word2); i++ {
		b += word2[i]
	}

	return a == b
}

func main() {
	word1 := []string{"a", "bc"}
	word2 := []string{"ab", "c"}

	fmt.Println(arrayStringsAreEqual(word1, word2))
}
