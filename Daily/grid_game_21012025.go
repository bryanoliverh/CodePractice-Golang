func gridGame(grid [][]int) int64 {
	n := len(grid[0])

	// Prefix sums for top and bottom rows, to quickly sum parts of the rows.
	topPrefixSum := make([]int, n+1)
	bottomPrefixSum := make([]int, n+1)

	for i := 0; i < n; i++ {
		topPrefixSum[i+1] = topPrefixSum[i] + grid[0][i]
		bottomPrefixSum[i+1] = bottomPrefixSum[i] + grid[1][i]
	}

	minSecondRobotPoints := int64(^uint64(0) >> 1) // Init with max int64 value

	for i := 0; i < n; i++ {
		// Points remaining at the top row if first robot drops down after column i
		topRemaining := topPrefixSum[n] - topPrefixSum[i+1]
		// Points collected up to column i in the bottom row (excludes the drop point)
		bottomCollected := bottomPrefixSum[i]

		// Second robot's potential points after the first robot switches
		secondRobotPoints := int64(max(topRemaining, bottomCollected))

		if secondRobotPoints < minSecondRobotPoints {
			minSecondRobotPoints = secondRobotPoints
		}
	}

	return minSecondRobotPoints
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}