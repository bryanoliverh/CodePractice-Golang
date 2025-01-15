

func minimizeXor(num1 int, num2 int) int {
	// Count the number of set bits in num2
	setBitsNum2 := bits.OnesCount(uint(num2))

	// Initialize the result to 0
	result := 0

	// First loop: try to match high-order bits of num1
	for i := 31; i >= 0 && setBitsNum2 > 0; i-- {
		// Check if the bit at position i in num1 is set
		if (num1 & (1 << i)) != 0 {
			// Set the bit at position i in result
			result |= (1 << i)
			// Decrease the count of set bits needed
			setBitsNum2--
		}
	}

	// Second loop: fill remaining set bits from low-order positions
	for i := 0; i < 32 && setBitsNum2 > 0; i++ {
		// Check if the bit at position i in result is not set
		if (result & (1 << i)) == 0 {
			// Set the bit at position i in result
			result |= (1 << i)
			// Decrease the count of set bits needed
			setBitsNum2--
		}
	}

	return result
}