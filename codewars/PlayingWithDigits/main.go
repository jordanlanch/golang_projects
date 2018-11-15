package main

import (
	"math"
)

//DigPow n pjajnas
func DigPow(n, p int) int {
	N := n
	var digits []int

	for n > 0 {
		digits = append(digits, n%10)
		n = n / 10
	}

	sum := 0
	for i := len(digits) - 1; i >= 0; i-- {
		sum = sum + int(math.Pow(float64(digits[i]), float64(p)))
		p = p + 1
	}

	if sum%N == 0 {
		return sum / N
	} else {
		return -1
	}

}
