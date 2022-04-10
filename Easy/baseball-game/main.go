package main

import (
	"fmt"
	"strconv"
)

func calPoints(ops []string) int {
	stack := make([]int, len(ops))
	top := 0

	for i := 0; i < len(ops); i++ {
		op := ops[i]

		if op == "+" {
			last1 := stack[top-1]
			last2 := stack[top-2]
			stack[top] = last1 + last2
			top++
		} else if op == "D" {
			last1 := stack[top-1]
			stack[top] = last1 * 2
			top++
		} else if op == "C" {
			top--
		} else {
			stack[top], _ = strconv.Atoi(op)
			top++
		}
	}

	points := 0

	for i := 0; i < top; i++ {
		points += stack[i]
	}

	return points
}

func main() {
	points := calPoints([]string{"5", "2", "C", "D", "+"})

	fmt.Printf("Result: %d", points)
}
