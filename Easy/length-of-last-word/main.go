package lengthoflastword

func lengthOfLastWord(s string) int {
	sl := len(s)

	lLastWord := 0

	for i := sl - 1; i >= 0; {
		for i >= 0 && string(s[i]) == " " {
			i--
		}

		if i < 0 {
			return 0
		}

		for i >= 0 && string(s[i]) != " " {
			i--
			lLastWord++
		}

		return lLastWord
	}

	return 0
}

func main() {
	lengthOfLastWord("   hello, world")
}
