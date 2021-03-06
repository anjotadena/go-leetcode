package romantointeger

func main(s string) int {
	if s == "" {
		return 0
	}

	symbols := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	symbols["I"] = 1
	symbols["V"] = 5
	symbols["X"] = 10
	symbols["L"] = 50
	symbols["C"] = 100
	symbols["D"] = 500
	symbols["M"] = 1000

	num, lastNum, total := 0, 0, 0

	sl := len(s)

	for i := 0; i < sl; i++ {
		char := s[(sl - (i + 1)):(sl - i)]

		num = symbols[char]

		if num < lastNum {
			total = total - num
		} else {
			total = total + num
		}

		lastNum = num
	}

	return total
}
