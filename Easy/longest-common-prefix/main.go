package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// assume prefix for first item
	p := strs[0]

	for i := 1; i < len(strs); i++ {
		if strings.Index(strs[i], p) != 0 {
			p = p[:len(p)-1]

			i--
		}
	}

	return p
}

func main() {
	prefix := longestCommonPrefix([]string{"flower", "flow", "flight"})

	fmt.Printf("Longest Prefix: \"%s\"\n", prefix)
}
