package main

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1

	for i <= j {
		// Swapping variables
		s[i], s[j] = s[j], s[i]

		//
		i++

		//
		j--
	}
}

func main() {
	reverseString([]byte("hello"))
}
