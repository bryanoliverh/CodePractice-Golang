package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var reversed int64

	// Handle negative input
	negative := false
	if x < 0 {
		negative = true
		x *= -1
	}

	// Reverse the digits
	for x > 0 {
		reversed = reversed*10 + int64(x%10)
		x /= 10
	}

	// Handle out of range values
	if reversed < math.MinInt32 || reversed > math.MaxInt32 {
		return 0
	}

	// Apply negative sign if necessary
	if negative {
		reversed *= -1
	}

	return int(reversed)
}

