package stringtointeger

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)

	if s == "" {
		return 0
	}

	pointer := 0
	fmt.Println(s)
	// positive or negative integer value
	sign := 1

	if s[pointer] == '+' {
		pointer += 1
	} else if s[pointer] == '-' {
		pointer += 1
		sign = -1
	}
	fmt.Println(pointer)
	fmt.Println(s)
	parsedValue := 0

	for pointer < len(s) {
		current := s[pointer] - '0'

		if !(current >= 0 && current <= 9) {
			break
		}

		if parsedValue > math.MaxInt32/10 || parsedValue == math.MaxInt32/10 && current > math.MaxInt32%10 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		parsedValue = (parsedValue * 10) + int(current)

		pointer++
	}

	return parsedValue * sign
}
