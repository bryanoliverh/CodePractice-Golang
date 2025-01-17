func doesValidArrayExist(derived []int) bool {
	n := len(derived)

	// Function to check if we can build a valid original with the initial value of original[0] = start
	canBuild := func(start int) bool {
		original := make([]int, n)
		original[0] = start

		for i := 1; i < n; i++ {
			original[i] = original[i-1] ^ derived[i-1]
		}

		// Check the last condition derived[n-1] = original[n-1] ⊕ original[0]
		return derived[n-1] == (original[n-1] ^ original[0])
	}

	// Check with original[0] = 0 or original[0] = 1
	return canBuild(0) || canBuild(1)
}