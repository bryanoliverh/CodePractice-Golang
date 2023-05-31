
func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    maxArea := 0

    for left < right {
        // Calculate the area between the lines
        area := min(height[left], height[right]) * (right - left)
        maxArea = max(maxArea, area)

        // Move the pointer with the smaller height inward
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }

    return maxArea
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
