func reverseWords(s string) string {
	ss := strings.Split(s, " ")
	answer := ""

	for i := 0; i < len(ss); i++ {
		answer += reverseString(ss[i])

		if i != len(ss)-1 {
			answer += " "
		}
	}
	return answer
}

func reverseString(s string) string {
	_s := []byte(s)
	l, r := 0, len(s)-1

	for l < r {
		_s[l], _s[r] = _s[r], _s[l]
		l++
		r--
	}

	return string(_s)
}

func main() {

}